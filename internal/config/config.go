package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Port                 string
	AllowedOrigins       []string
	HTTPReadTimeout      time.Duration
	HTTPWriteTimeout     time.Duration
	HTTPIdleTimeout      time.Duration
	OllamaHost           string
	OllamaBin            string
	OllamaKeepAlive      string
	OllamaContextLength  string
	OllamaKVCacheType    string
	OllamaFlashAttention string
	AssistantModel       string
	SiteBaseURL          string
	SiteSeedPaths        []string
	SiteRequestTimeout   time.Duration
	SiteRefreshInterval  time.Duration
	SiteMaxPages         int
	SiteMaxContentChars  int
	SiteUserAgent        string
}

func Load() Config {
	loadEnvFiles(envFileCandidates()...)

	return Config{
		Port:                 envOrDefault("PORT", "10123"),
		AllowedOrigins:       envOrDefaultCSV("CORS_ALLOWED_ORIGINS", []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://winapps.io", "https://www.winapps.io"}),
		HTTPReadTimeout:      envOrDefaultDuration("HTTP_READ_TIMEOUT", 30*time.Second),
		HTTPWriteTimeout:     envOrDefaultDuration("HTTP_WRITE_TIMEOUT", 10*time.Minute),
		HTTPIdleTimeout:      envOrDefaultDuration("HTTP_IDLE_TIMEOUT", 2*time.Minute),
		OllamaHost:           envOrDefault("OLLAMA_HOST", "http://127.0.0.1:11434"),
		OllamaBin:            envOrDefault("OLLAMA_BIN", "ollama"),
		OllamaKeepAlive:      envOrDefault("OLLAMA_KEEP_ALIVE", "0m"),
		OllamaContextLength:  envOrDefault("OLLAMA_CONTEXT_LENGTH", "16384"),
		OllamaKVCacheType:    envOrDefault("OLLAMA_KV_CACHE_TYPE", "q8_0"),
		OllamaFlashAttention: envOrDefault("OLLAMA_FLASH_ATTENTION", "1"),
		AssistantModel:       envOrDefault("OLLAMA_MODEL", "deepseek-r1:8b-llama-distill-q8_0"),
		SiteBaseURL:          strings.TrimRight(envOrDefault("SITE_BASE_URL", "https://winapps.io"), "/"),
		SiteSeedPaths:        envOrDefaultCSV("SITE_SEED_PATHS", nil),
		SiteRequestTimeout:   envOrDefaultDuration("SITE_REQUEST_TIMEOUT", 10*time.Second),
		SiteRefreshInterval:  envOrDefaultDuration("SITE_REFRESH_INTERVAL", 30*time.Minute),
		SiteMaxPages:         envOrDefaultInt("SITE_MAX_PAGES", 48),
		SiteMaxContentChars:  envOrDefaultInt("SITE_MAX_CONTENT_CHARS", 5000),
		SiteUserAgent:        envOrDefault("SITE_USER_AGENT", "WinnyAI/1.0 (+https://winapps.io)"),
	}
}

func envFileCandidates() []string {
	roots := []string{"."}
	if exePath, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exePath)
		roots = append(roots, exeDir, filepath.Dir(exeDir))
	}

	relativePaths := []string{
		".env",
		".env.local",
	}

	seen := make(map[string]struct{})
	candidates := make([]string, 0, len(roots)*len(relativePaths))
	for _, root := range roots {
		for _, relPath := range relativePaths {
			candidate := relPath
			if root != "." {
				candidate = filepath.Join(root, relPath)
			}
			if _, ok := seen[candidate]; ok {
				continue
			}
			seen[candidate] = struct{}{}
			candidates = append(candidates, candidate)
		}
	}
	return candidates
}

func envOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}

func envOrDefaultCSV(key string, fallback []string) []string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return append([]string(nil), fallback...)
	}

	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			out = append(out, part)
		}
	}
	if len(out) == 0 {
		return append([]string(nil), fallback...)
	}
	return out
}

func envOrDefaultInt(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	var parsed int
	if _, err := fmt.Sscan(value, &parsed); err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func envOrDefaultDuration(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	parsed, err := time.ParseDuration(value)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func loadEnvFiles(paths ...string) {
	for _, path := range paths {
		loadEnvFile(path)
	}
}

func loadEnvFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimPrefix(line, "export ")
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		if key == "" || os.Getenv(key) != "" {
			continue
		}
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		_ = os.Setenv(key, value)
	}
}
