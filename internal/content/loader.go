package content

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"io.winapps/internal/config"
	"io.winapps/internal/domain"
)

func LoadSitePages(cfg config.Config) ([]domain.SitePage, error) {
	rootDir, err := resolveContentRoot(cfg)
	if err != nil {
		return nil, err
	}

	pages := make([]domain.SitePage, 0, 48)
	err = filepath.WalkDir(rootDir, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		if strings.ToLower(filepath.Ext(path)) != ".md" {
			return nil
		}

		relativePath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		bytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		markdown := strings.TrimSpace(string(bytes))
		if markdown == "" {
			return nil
		}

		route := routeFromRelativePath(relativePath)
		title := extractTitle(markdown, relativePath)
		plainText := markdownToText(markdown)
		url := ""
		if strings.HasPrefix(route, "/") && !strings.HasPrefix(route, "/_docs/") {
			url = strings.TrimRight(cfg.SiteBaseURL, "/")
			if route != "/" {
				url += route
			} else {
				url += "/"
			}
		}

		pages = append(pages, domain.SitePage{
			URL:        url,
			Path:       route,
			SourcePath: filepath.ToSlash(filepath.Join("internal", "content", relativePath)),
			Title:      title,
			Summary:    truncateString(plainText, 900),
			Content:    truncateString(markdown, 5000),
			Tokens:     tokenize(strings.Join([]string{title, route, plainText}, " ")),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(pages) == 0 {
		return nil, fmt.Errorf("no markdown pages found under %s", rootDir)
	}

	sort.SliceStable(pages, func(i, j int) bool {
		return pages[i].SourcePath < pages[j].SourcePath
	})
	return pages, nil
}

func resolveContentRoot(cfg config.Config) (string, error) {
	candidates := []string{}
	if strings.TrimSpace(cfg.ContentDir) != "" {
		candidates = append(candidates, cfg.ContentDir)
	}
	candidates = append(candidates, "./internal/content")

	if exePath, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exePath)
		candidates = append(candidates,
			filepath.Join(exeDir, "internal", "content"),
			filepath.Join(filepath.Dir(exeDir), "internal", "content"),
		)
	}

	seen := map[string]struct{}{}
	for _, candidate := range candidates {
		candidate = strings.TrimSpace(candidate)
		if candidate == "" {
			continue
		}
		absolute, err := filepath.Abs(candidate)
		if err != nil {
			continue
		}
		if _, ok := seen[absolute]; ok {
			continue
		}
		seen[absolute] = struct{}{}
		info, err := os.Stat(absolute)
		if err == nil && info.IsDir() {
			return absolute, nil
		}
	}

	return "", errors.New("content directory not found")
}

func routeFromRelativePath(relativePath string) string {
	relativePath = filepath.ToSlash(relativePath)
	if relativePath == "page.md" {
		return "/"
	}
	if strings.HasSuffix(relativePath, "/page.md") {
		route := strings.TrimSuffix(relativePath, "/page.md")
		if route == "" {
			return "/"
		}
		return "/" + route
	}
	return "/_docs/" + strings.TrimSuffix(relativePath, ".md")
}

func extractTitle(markdown, relativePath string) string {
	for _, line := range strings.Split(markdown, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}

	base := filepath.Base(relativePath)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func markdownToText(markdown string) string {
	replacer := strings.NewReplacer(
		"```", " ",
		"`", " ",
		"#", " ",
		"*", " ",
		"_", " ",
		">", " ",
		"|", " ",
		"[", " ",
		"]", " ",
		"(", " ",
		")", " ",
	)
	markdown = replacer.Replace(markdown)
	return normalizeWhitespace(markdown)
}

func tokenize(value string) map[string]struct{} {
	value = strings.ToLower(value)
	fields := strings.FieldsFunc(value, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= '0' && r <= '9')
	})
	if len(fields) == 0 {
		return nil
	}

	stopWords := map[string]struct{}{
		"the": {}, "and": {}, "for": {}, "that": {}, "with": {}, "this": {},
		"from": {}, "your": {}, "have": {}, "will": {}, "into": {}, "about": {},
		"what": {}, "when": {}, "where": {}, "which": {}, "their": {}, "they": {},
		"them": {}, "just": {}, "here": {}, "page": {}, "site": {},
	}

	tokens := make(map[string]struct{}, len(fields))
	for _, field := range fields {
		if len(field) < 2 {
			continue
		}
		if _, skip := stopWords[field]; skip {
			continue
		}
		tokens[field] = struct{}{}
	}
	return tokens
}

func normalizeWhitespace(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), " ")
}

func truncateString(value string, limit int) string {
	if limit <= 0 || len(value) <= limit {
		return value
	}
	if limit <= 3 {
		return value[:limit]
	}
	return strings.TrimSpace(value[:limit-3]) + "..."
}
