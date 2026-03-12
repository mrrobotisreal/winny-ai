package runner

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"io.winapps/internal/config"
	"io.winapps/internal/domain"
)

type OllamaCLIRunner struct{ cfg config.Config }

func NewOllamaCLIRunner(cfg config.Config) *OllamaCLIRunner {
	return &OllamaCLIRunner{cfg: cfg}
}

func (r *OllamaCLIRunner) Stop(ctx context.Context, model string) error {
	if strings.TrimSpace(model) == "" {
		return nil
	}

	cmd := exec.CommandContext(ctx, r.cfg.OllamaBin, "stop", model)
	cmd.Env = runnerEnv(r.cfg, domain.ChatSettings{})
	output, err := cmd.CombinedOutput()
	if err != nil {
		text := strings.TrimSpace(string(output))
		if strings.Contains(strings.ToLower(text), "not loaded") {
			return nil
		}
		return fmt.Errorf("ollama stop: %w (%s)", err, text)
	}
	return nil
}

func (r *OllamaCLIRunner) Stream(ctx context.Context, settings domain.ChatSettings, history []domain.Message, onChunk func(domain.AssistantDelta) error) (domain.AssistantResponse, error) {
	prompt := buildPrompt(history)
	stdin := new(bytes.Buffer)
	if settings.SystemPrompt != "" {
		fmt.Fprintf(stdin, "/set system %s\n", settings.SystemPrompt)
	}
	fmt.Fprintln(stdin, "/set nothink")
	fmt.Fprintf(stdin, "/set parameter temperature %v\n", settings.Temperature)
	fmt.Fprintf(stdin, "/set parameter top_p %v\n", settings.TopP)
	fmt.Fprintf(stdin, "/set parameter top_k %d\n", settings.TopK)
	fmt.Fprintf(stdin, "/set parameter num_ctx %d\n", settings.NumCtx)
	fmt.Fprintf(stdin, "/set parameter num_predict %d\n", settings.NumPredict)
	fmt.Fprintf(stdin, "/set parameter repeat_penalty %v\n", settings.RepeatPenalty)
	fmt.Fprintf(stdin, "%s\n", prompt)

	cmd := exec.CommandContext(ctx, r.cfg.OllamaBin, "run", "--think=false", settings.TextModel)
	cmd.Stdin = stdin
	cmd.Env = runnerEnv(r.cfg, settings)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return domain.AssistantResponse{}, fmt.Errorf("stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return domain.AssistantResponse{}, fmt.Errorf("stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return domain.AssistantResponse{}, fmt.Errorf("start ollama: %w", err)
	}

	var stderrText strings.Builder
	var stderrMu sync.Mutex
	var stderrWG sync.WaitGroup
	stderrWG.Add(1)
	go func() {
		defer stderrWG.Done()
		if bytes, err := io.ReadAll(stderr); err == nil && len(bytes) > 0 {
			stderrMu.Lock()
			stderrText.Write(bytes)
			stderrMu.Unlock()
		}
	}()

	var content strings.Builder
	filter := newThinkTagFilter(func(chunk string) error {
		if chunk == "" {
			return nil
		}
		content.WriteString(chunk)
		if onChunk != nil {
			return onChunk(domain.AssistantDelta{Content: chunk, Model: settings.TextModel})
		}
		return nil
	})

	buf := make([]byte, 2048)
	for {
		n, readErr := stdout.Read(buf)
		if n > 0 {
			chunk := sanitizeCLIChunk(string(buf[:n]))
			if err := filter.Push(chunk); err != nil {
				_ = cmd.Process.Kill()
				stderrWG.Wait()
				return domain.AssistantResponse{}, err
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			_ = cmd.Process.Kill()
			stderrWG.Wait()
			return domain.AssistantResponse{}, fmt.Errorf("read ollama output: %w", readErr)
		}
	}

	if err := filter.Flush(); err != nil {
		_ = cmd.Process.Kill()
		stderrWG.Wait()
		return domain.AssistantResponse{}, err
	}

	waitErr := cmd.Wait()
	stderrWG.Wait()
	if waitErr != nil {
		stderrMu.Lock()
		errText := strings.TrimSpace(stderrText.String())
		stderrMu.Unlock()
		return domain.AssistantResponse{}, fmt.Errorf("ollama run: %w (%s)", waitErr, errText)
	}

	resp := domain.AssistantResponse{
		Content: strings.TrimSpace(content.String()),
		Model:   settings.TextModel,
	}
	if onChunk != nil {
		if err := onChunk(domain.AssistantDelta{Model: resp.Model, Done: true}); err != nil {
			return domain.AssistantResponse{}, err
		}
	}
	return resp, nil
}

func buildPrompt(messages []domain.Message) string {
	var b strings.Builder
	for _, message := range messages {
		role := strings.ToUpper(strings.TrimSpace(message.Role))
		if role == "" {
			role = "USER"
		}
		b.WriteString("[" + role + "]\n")
		b.WriteString(strings.TrimSpace(message.Content))
		b.WriteString("\n\n")
	}
	return strings.TrimSpace(b.String())
}

func runnerEnv(cfg config.Config, settings domain.ChatSettings) []string {
	contextLength := cfg.OllamaContextLength
	if settings.NumCtx > 0 {
		contextLength = strconv.Itoa(settings.NumCtx)
	}

	keepAlive := cfg.OllamaKeepAlive
	if strings.TrimSpace(settings.KeepAlive) != "" {
		keepAlive = settings.KeepAlive
	}

	env := os.Environ()
	env = append(env,
		"OLLAMA_HOST="+cfg.OllamaHost,
		"OLLAMA_KEEP_ALIVE="+keepAlive,
		"OLLAMA_CONTEXT_LENGTH="+contextLength,
		"OLLAMA_KV_CACHE_TYPE="+cfg.OllamaKVCacheType,
		"OLLAMA_FLASH_ATTENTION="+cfg.OllamaFlashAttention,
		"NO_COLOR=1",
		"TERM=dumb",
	)
	return env
}

func sanitizeCLIChunk(chunk string) string {
	chunk = strings.ReplaceAll(chunk, "\r", "")
	return chunk
}

type thinkTagFilter struct {
	inThink bool
	carry   string
	emit    func(string) error
}

func newThinkTagFilter(emit func(string) error) *thinkTagFilter {
	return &thinkTagFilter{emit: emit}
}

func (f *thinkTagFilter) Push(chunk string) error {
	data := f.carry + chunk
	f.carry = ""

	for len(data) > 0 {
		if f.inThink {
			if idx := strings.Index(data, "</think>"); idx >= 0 {
				data = data[idx+len("</think>"):]
				f.inThink = false
				continue
			}
			f.carry = partialTagSuffix(data, "</think>")
			return nil
		}

		if idx := strings.Index(data, "<think>"); idx >= 0 {
			if idx > 0 {
				if err := f.emit(data[:idx]); err != nil {
					return err
				}
			}
			data = data[idx+len("<think>"):]
			f.inThink = true
			continue
		}

		tail := partialTagSuffix(data, "<think>")
		emitPart := data[:len(data)-len(tail)]
		if emitPart != "" {
			if err := f.emit(emitPart); err != nil {
				return err
			}
		}
		f.carry = tail
		return nil
	}

	return nil
}

func (f *thinkTagFilter) Flush() error {
	if !f.inThink && f.carry != "" {
		if err := f.emit(f.carry); err != nil {
			return err
		}
	}
	f.carry = ""
	return nil
}

func partialTagSuffix(data, tag string) string {
	max := len(tag) - 1
	if max > len(data) {
		max = len(data)
	}
	for i := max; i > 0; i-- {
		if strings.HasSuffix(data, tag[:i]) {
			return data[len(data)-i:]
		}
	}
	return ""
}
