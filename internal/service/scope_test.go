package service

import (
	"strings"
	"testing"

	"io.winapps/internal/catalog"
	"io.winapps/internal/config"
	"io.winapps/internal/content"
	"io.winapps/internal/domain"
)

func loadPages(t *testing.T) []domain.SitePage {
	t.Helper()
	pages, err := content.LoadSitePages(config.Config{ContentDir: "../content", SiteBaseURL: "https://www.winapps.io"})
	if err != nil {
		t.Fatalf("load content: %v", err)
	}
	return pages
}

func ask(text string) domain.ChatRequest {
	return domain.ChatRequest{Messages: []domain.Message{{Role: "user", Content: text}}}
}

func TestScopeGate(t *testing.T) {
	pages := loadPages(t)
	terms := siteTerms(pages)
	cases := map[string]bool{
		"What is Сумрак?":                      true,
		"Is Mitch still at Mirror?":            true,
		"How much revenue did prebidding add?": true,
		"What does Darkroom do?":               true,
		"Tell me about ODIN":                   true,
		"What's the weather?":                  false,
		"What's your favorite pizza?":          false,
	}
	for question, want := range cases {
		req := ask(question)
		got := isWebsiteQuestion(req, rankRelevantPages(req, pages), terms, "", "")
		if got != want {
			t.Errorf("%q: in scope = %v, want %v", question, got, want)
		}
	}
}

func TestPromptCarriesFactsNotHardcodedRoles(t *testing.T) {
	pages := loadPages(t)
	facts, ok := content.FactsPage(pages)
	if !ok {
		t.Fatal("facts.md not loaded")
	}
	req := ask("Is Mitch still at Mirror?")
	prompt := catalog.BuildSystemPrompt(req, rankRelevantPages(req, pages), facts.Content, "", "")
	if !strings.Contains(prompt, "ended in May 2026") {
		t.Error("prompt is missing the Mirror end date from facts.md")
	}
	if strings.Contains(prompt, "primary mission outside of work") {
		t.Error("prompt still contains the old hard-coded CreaTV mission line")
	}
	for _, page := range rankRelevantPages(req, pages) {
		if page.Path == "/_docs/facts" {
			t.Error("facts doc must not be ranked as a page")
		}
	}
}
