package content

import (
	"testing"

	"io.winapps/internal/config"
)

func TestLoadGeneratedContent(t *testing.T) {
	pages, err := LoadSitePages(config.Config{ContentDir: "."})
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	byPath := map[string]bool{}
	for _, p := range pages {
		byPath[p.Path] = true
	}
	for _, want := range []string{"/", "/projects/sumrak", "/portfolio/people-inc", "/portfolio/mirror", "/_docs/facts"} {
		if !byPath[want] {
			t.Errorf("missing %s", want)
		}
	}
	for _, dead := range []string{"/projects/rss-today", "/projects/media-manipulator-pro", "/portfolio/double-raven"} {
		if byPath[dead] {
			t.Errorf("stale route %s still present", dead)
		}
	}
	if _, ok := FactsPage(pages); !ok {
		t.Error("facts page not found")
	}
}

func TestExtractKeywords(t *testing.T) {
	md := "# T\n\n## Search Keywords\n`Sumrak`, `Сумрак`, `sumrak`, `Darkroom`\n\n## Important Links\n- `not-a-keyword`\n"
	got := extractKeywords(md)
	want := []string{"sumrak", "сумрак", "darkroom"}
	if len(got) != len(want) {
		t.Fatalf("got %v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v want %v", got, want)
		}
	}
}
