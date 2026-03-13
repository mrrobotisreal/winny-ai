package service

import (
	"context"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"io.winapps/internal/catalog"
	"io.winapps/internal/config"
	"io.winapps/internal/content"
	"io.winapps/internal/domain"
)

var (
	ErrInvalidRequest = errors.New("at least one non-empty message is required")

	titleRe      = regexp.MustCompile(`(?is)<title[^>]*>(.*?)</title>`)
	headRe       = regexp.MustCompile(`(?is)<head[^>]*>.*?</head>`)
	scriptRe     = regexp.MustCompile(`(?is)<script[^>]*>.*?</script>`)
	styleRe      = regexp.MustCompile(`(?is)<style[^>]*>.*?</style>`)
	noscriptRe   = regexp.MustCompile(`(?is)<noscript[^>]*>.*?</noscript>`)
	svgRe        = regexp.MustCompile(`(?is)<svg[^>]*>.*?</svg>`)
	commentRe    = regexp.MustCompile(`(?s)<!--.*?-->`)
	tagRe        = regexp.MustCompile(`(?s)<[^>]+>`)
	whitespaceRe = regexp.MustCompile(`\s+`)
	tokenRe      = regexp.MustCompile(`[\p{L}\p{N}]+`)

	languageRequestPhrases = map[string][]string{
		"English": {
			"respond in english",
			"answer in english",
			"reply in english",
			"write in english",
		},
		"Russian": {
			"respond in russian",
			"answer in russian",
			"reply in russian",
			"write in russian",
			"на русском",
			"по-русски",
			"ответь по-русски",
			"ответьте по-русски",
		},
		"Ukrainian": {
			"respond in ukrainian",
			"answer in ukrainian",
			"reply in ukrainian",
			"write in ukrainian",
			"українською",
			"відповідай українською",
			"відповідь українською",
		},
		"Hebrew": {
			"respond in hebrew",
			"answer in hebrew",
			"reply in hebrew",
			"write in hebrew",
			"ענה בעברית",
			"בעברית",
		},
		"Modern Standard Arabic": {
			"respond in arabic",
			"answer in arabic",
			"reply in arabic",
			"write in arabic",
			"أجب بالعربية",
			"بالعربية",
		},
		"Spanish": {
			"respond in spanish",
			"answer in spanish",
			"reply in spanish",
			"write in spanish",
			"responde en español",
			"en español por favor",
		},
		"Vietnamese": {
			"respond in vietnamese",
			"answer in vietnamese",
			"reply in vietnamese",
			"write in vietnamese",
			"trả lời bằng tiếng việt",
			"bằng tiếng việt",
		},
		"Mandarin Chinese": {
			"respond in chinese",
			"answer in chinese",
			"reply in chinese",
			"write in chinese",
			"请用中文",
			"用中文回答",
		},
		"French": {
			"respond in french",
			"answer in french",
			"reply in french",
			"write in french",
			"réponds en français",
			"en français",
		},
		"German": {
			"respond in german",
			"answer in german",
			"reply in german",
			"write in german",
			"antworte auf deutsch",
			"auf deutsch",
		},
		"Portuguese": {
			"respond in portuguese",
			"answer in portuguese",
			"reply in portuguese",
			"write in portuguese",
			"responda em português",
			"em português",
		},
		"Italian": {
			"respond in italian",
			"answer in italian",
			"reply in italian",
			"write in italian",
			"rispondi in italiano",
			"in italiano",
		},
		"Japanese": {
			"respond in japanese",
			"answer in japanese",
			"reply in japanese",
			"write in japanese",
			"日本語で答えて",
			"日本語で",
		},
	}

	websiteTerms = []string{
		"winny",
		"winapps",
		"winapps.io",
		"website",
		"site",
		"page",
		"mitch",
		"mitchell",
		"mitchell wintrow",
		"project",
		"portfolio",
		"blog",
		"contact",
		"about",
		"download",
		"support",
		"documentation",
		"privacy",
		"terms",
		"journey app",
		"journey",
		"voizy",
		"creatv",
		"creatv.io",
		"mirror",
		"mirror inc",
		"trymirror.ai",
		"media manipulator",
		"finance flow",
		"quantum river",
		"rss today",
		"text formatter",
		"mp3 drive player",
		"people inc",
		"mobile ad revenue",
		"vmware",
		"amazon connect",
		"double raven",
		"alina",
		"ukraine",
		"drone",
		"dji mini 4 pro",
		"ai server",
		"home server",
		"full tower server",
		"language learning",
		"russian",
		"ukrainian",
		"hebrew",
		"farsi",
		"modern standard arabic",
		"mandarin chinese",
		"spanish",
		"vietnamese",
		"сайт",
		"страница",
		"митч",
		"митчелл",
		"портфолио",
		"проект",
		"контакт",
		"блог",
		"сторінка",
		"мітч",
		"мітчелл",
		"портфоліо",
		"проєкт",
		"контакти",
		"אתר",
		"עמוד",
		"מיטש",
		"מיטשל",
		"פרויקט",
		"פורטפוליו",
		"בלוג",
		"יצירת קשר",
		"موقع",
		"صفحة",
		"ميتش",
		"ميتشل",
		"مشروع",
		"مدونة",
		"اتصال",
		"sitio",
		"página",
		"portafolio",
		"proyecto",
		"contacto",
		"trang",
		"trang web",
		"dự án",
		"liên hệ",
		"tải xuống",
		"网站",
		"页面",
		"项目",
		"作品集",
		"联系",
		"下载",
		"米切尔",
		"ce site",
		"cette page",
		"projet",
		"diese seite",
		"diese website",
		"projekt",
		"kontakt",
		"este site",
		"esta página",
		"projeto",
		"contato",
		"questa pagina",
		"questo sito",
		"progetto",
		"contatto",
		"このページ",
		"このサイト",
		"プロジェクト",
		"連絡先",
		"ミッチェル",
	}

	pageScopedPhrases = []string{
		"this page",
		"current page",
		"on this site",
		"on this page",
		"here on the site",
		"what does this page",
		"tell me about this page",
		"эта страница",
		"на этом сайте",
		"на этой странице",
		"ця сторінка",
		"на цьому сайті",
		"на цій сторінці",
		"העמוד הזה",
		"באתר הזה",
		"בדף הזה",
		"هذه الصفحة",
		"هذا الموقع",
		"في هذه الصفحة",
		"esta página",
		"este sitio",
		"en esta página",
		"trang này",
		"trên trang này",
		"这个页面",
		"这个网站",
		"在这个页面",
		"cette page",
		"ce site",
		"auf dieser seite",
		"este site",
		"questa pagina",
		"questo sito",
		"このページ",
		"このサイト",
	}
)

type AssistantRunner interface {
	Stream(context.Context, domain.ChatSettings, []domain.Message, func(domain.AssistantDelta) error) (domain.AssistantResponse, error)
	Stop(context.Context, string) error
}

type ChatService struct {
	runner          AssistantRunner
	cfg             config.Config
	defaultSettings domain.ChatSettings
	siteCatalog     *siteCatalog
}

func NewChatService(runner AssistantRunner, cfg config.Config) *ChatService {
	settings := catalog.DefaultChatSettings()
	settings.TextModel = cfg.AssistantModel
	settings.KeepAlive = cfg.OllamaKeepAlive

	service := &ChatService{
		runner:          runner,
		cfg:             cfg,
		defaultSettings: settings,
		siteCatalog:     newSiteCatalog(cfg),
	}

	go service.warmSiteCatalog()
	return service
}

func (s *ChatService) StreamChat(ctx context.Context, request domain.ChatRequest, onEvent func(domain.StreamEvent) error) error {
	messages := sanitizeMessages(request.Messages)
	if len(messages) == 0 {
		return ErrInvalidRequest
	}

	request.Messages = messages

	pages, pagesErr := s.siteCatalog.pages(ctx)
	relevantPages := rankRelevantPages(request, pages)
	preferredLanguage := inferPreferredResponseLanguage(request)
	browserLanguageHint := supportedBrowserLanguageHint(request.BrowserLanguage)
	if !isWebsiteQuestion(request, relevantPages, preferredLanguage, browserLanguageHint) {
		return s.streamStaticReply(onEvent, catalog.OutOfScopeReply())
	}

	if err := onEvent(domain.StreamEvent{
		Type:      "meta",
		Assistant: catalog.AssistantName,
		Model:     s.defaultSettings.TextModel,
	}); err != nil {
		return err
	}

	settings := s.defaultSettings
	settings.SystemPrompt = catalog.BuildSystemPrompt(
		request,
		relevantPages,
		preferredLanguage,
		browserLanguageHint,
	)

	_, err := s.runner.Stream(ctx, settings, messages, func(delta domain.AssistantDelta) error {
		if delta.Done {
			return onEvent(domain.StreamEvent{
				Type:  "done",
				Model: delta.Model,
			})
		}
		if delta.Content == "" {
			return nil
		}
		return onEvent(domain.StreamEvent{
			Type:  "delta",
			Model: delta.Model,
			Delta: delta.Content,
		})
	})
	if err == nil {
		return nil
	}

	if pagesErr != nil {
		return fmt.Errorf("site context unavailable: %w; assistant run failed: %w", pagesErr, err)
	}
	return err
}

func (s *ChatService) Reset(ctx context.Context) error {
	return s.runner.Stop(ctx, s.defaultSettings.TextModel)
}

func (s *ChatService) warmSiteCatalog() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	_, _ = s.siteCatalog.pages(ctx)
}

func (s *ChatService) streamStaticReply(onEvent func(domain.StreamEvent) error, reply string) error {
	if err := onEvent(domain.StreamEvent{
		Type:      "meta",
		Assistant: catalog.AssistantName,
		Model:     s.defaultSettings.TextModel,
	}); err != nil {
		return err
	}
	if err := onEvent(domain.StreamEvent{
		Type:  "delta",
		Model: s.defaultSettings.TextModel,
		Delta: reply,
	}); err != nil {
		return err
	}
	return onEvent(domain.StreamEvent{
		Type:  "done",
		Model: s.defaultSettings.TextModel,
	})
}

func sanitizeMessages(messages []domain.Message) []domain.Message {
	if len(messages) == 0 {
		return nil
	}

	const maxMessages = 12
	out := make([]domain.Message, 0, minInt(len(messages), maxMessages))
	for _, message := range messages {
		content := strings.TrimSpace(message.Content)
		if content == "" {
			continue
		}
		if len(content) > 4000 {
			content = content[:4000]
		}

		role := strings.ToLower(strings.TrimSpace(message.Role))
		switch role {
		case "assistant", "system":
			role = "assistant"
		default:
			role = "user"
		}

		out = append(out, domain.Message{
			Role:    role,
			Content: content,
		})
	}

	if len(out) > maxMessages {
		out = out[len(out)-maxMessages:]
	}
	return out
}

func isWebsiteQuestion(request domain.ChatRequest, relevant []domain.SitePage, preferredLanguage, browserLanguageHint string) bool {
	message := strings.TrimSpace(lastUserMessage(request.Messages))
	if message == "" {
		return false
	}
	messageLower := strings.ToLower(message)

	if containsAnySubstring(messageLower, websiteTerms) {
		return true
	}

	if len(relevant) > 0 {
		queryTokens := tokenize(messageLower)
		for _, page := range relevant {
			for token := range queryTokens {
				if _, ok := page.Tokens[token]; ok {
					return true
				}
			}
		}
	}

	if containsAnySubstring(messageLower, pageScopedPhrases) {
		return true
	}

	currentPath := normalizedPagePath(request.PagePath, request.PageURL)
	if currentPath != "" && hasCurrentPageMatch(relevant, currentPath) {
		if preferredLanguage != "" || browserLanguageHint != "" || usesSupportedNonLatinScript(message) {
			return true
		}
	}

	return false
}

func lastUserMessage(messages []domain.Message) string {
	for i := len(messages) - 1; i >= 0; i-- {
		if strings.EqualFold(messages[i].Role, "user") {
			return messages[i].Content
		}
	}
	return ""
}

func inferPreferredResponseLanguage(request domain.ChatRequest) string {
	message := strings.TrimSpace(lastUserMessage(request.Messages))
	if message == "" {
		return ""
	}
	if language := detectExplicitResponseLanguage(message); language != "" {
		return language
	}
	return detectLanguageByScript(message)
}

func supportedBrowserLanguageHint(browserLanguage string) string {
	value := strings.ToLower(strings.TrimSpace(browserLanguage))
	switch {
	case strings.HasPrefix(value, "ru"):
		return "Russian"
	case strings.HasPrefix(value, "uk"):
		return "Ukrainian"
	case strings.HasPrefix(value, "he"), strings.HasPrefix(value, "iw"):
		return "Hebrew"
	case strings.HasPrefix(value, "ar"):
		return "Modern Standard Arabic"
	case strings.HasPrefix(value, "es"):
		return "Spanish"
	case strings.HasPrefix(value, "vi"):
		return "Vietnamese"
	case strings.HasPrefix(value, "zh"):
		return "Mandarin Chinese"
	case strings.HasPrefix(value, "fr"):
		return "French"
	case strings.HasPrefix(value, "de"):
		return "German"
	case strings.HasPrefix(value, "pt"):
		return "Portuguese"
	case strings.HasPrefix(value, "it"):
		return "Italian"
	case strings.HasPrefix(value, "ja"):
		return "Japanese"
	default:
		return ""
	}
}

func detectExplicitResponseLanguage(message string) string {
	message = strings.ToLower(strings.TrimSpace(message))
	for language, phrases := range languageRequestPhrases {
		for _, phrase := range phrases {
			if strings.Contains(message, phrase) {
				return language
			}
		}
	}
	return ""
}

func detectLanguageByScript(message string) string {
	hasCyrillic := false
	hasUkrainianMarker := false
	hasHebrew := false
	hasArabic := false
	hasHan := false

	for _, r := range message {
		switch {
		case unicode.In(r, unicode.Hiragana, unicode.Katakana):
			return "Japanese"
		case unicode.In(r, unicode.Hebrew):
			hasHebrew = true
		case unicode.In(r, unicode.Arabic):
			hasArabic = true
		case unicode.In(r, unicode.Cyrillic):
			hasCyrillic = true
			if strings.ContainsRune("іїєґІЇЄҐ", r) {
				hasUkrainianMarker = true
			}
		case unicode.In(r, unicode.Han):
			hasHan = true
		}
	}

	switch {
	case hasHebrew:
		return "Hebrew"
	case hasArabic:
		return "Modern Standard Arabic"
	case hasCyrillic && hasUkrainianMarker:
		return "Ukrainian"
	case hasCyrillic:
		return "Russian"
	case hasHan:
		return "Mandarin Chinese"
	default:
		return ""
	}
}

func containsAnySubstring(message string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(message, term) {
			return true
		}
	}
	return false
}

func hasCurrentPageMatch(relevant []domain.SitePage, currentPath string) bool {
	for _, page := range relevant {
		if samePath(page.Path, currentPath) {
			return true
		}
	}
	return false
}

func usesSupportedNonLatinScript(message string) bool {
	return detectLanguageByScript(message) != ""
}

func rankRelevantPages(request domain.ChatRequest, pages []domain.SitePage) []domain.SitePage {
	if len(pages) == 0 {
		return nil
	}

	type scoredPage struct {
		page  domain.SitePage
		score int
	}

	currentPath := normalizedPagePath(request.PagePath, request.PageURL)
	queryTokens := tokenize(strings.Join([]string{
		lastUserMessage(request.Messages),
		request.PageTitle,
		currentPath,
	}, " "))

	scored := make([]scoredPage, 0, len(pages))
	for _, page := range pages {
		score := 0
		if currentPath != "" && samePath(page.Path, currentPath) {
			score += 100
		}

		titleLower := strings.ToLower(page.Title)
		pathLower := strings.ToLower(page.Path)
		for token := range queryTokens {
			if _, ok := page.Tokens[token]; ok {
				score += 4
			}
			if strings.Contains(titleLower, token) {
				score += 6
			}
			if strings.Contains(pathLower, token) {
				score += 8
			}
		}

		if score > 0 {
			scored = append(scored, scoredPage{page: page, score: score})
		}
	}

	sort.SliceStable(scored, func(i, j int) bool {
		if scored[i].score == scored[j].score {
			return scored[i].page.Path < scored[j].page.Path
		}
		return scored[i].score > scored[j].score
	})

	limit := minInt(4, len(scored))
	out := make([]domain.SitePage, 0, limit)
	for i := 0; i < limit; i++ {
		out = append(out, scored[i].page)
	}
	return out
}

func samePath(left, right string) bool {
	return strings.TrimRight(left, "/") == strings.TrimRight(right, "/")
}

func normalizedPagePath(pagePath, pageURL string) string {
	pagePath = strings.TrimSpace(pagePath)
	if pagePath != "" {
		if !strings.HasPrefix(pagePath, "/") {
			pagePath = "/" + pagePath
		}
		return pagePath
	}

	if pageURL == "" {
		return ""
	}
	parsed, err := url.Parse(pageURL)
	if err != nil {
		return ""
	}
	if parsed.Path == "" {
		return "/"
	}
	return parsed.Path
}

type siteCatalog struct {
	cfg    config.Config
	client *http.Client

	loadMu sync.Mutex

	mu         sync.RWMutex
	pagesCache []domain.SitePage
	loadedAt   time.Time
}

func newSiteCatalog(cfg config.Config) *siteCatalog {
	return &siteCatalog{
		cfg: cfg,
		client: &http.Client{
			Timeout: cfg.SiteRequestTimeout,
		},
	}
}

func (c *siteCatalog) pages(ctx context.Context) ([]domain.SitePage, error) {
	c.mu.RLock()
	pages := append([]domain.SitePage(nil), c.pagesCache...)
	loadedAt := c.loadedAt
	c.mu.RUnlock()

	if len(pages) > 0 && time.Since(loadedAt) < c.cfg.SiteRefreshInterval {
		return pages, nil
	}

	c.loadMu.Lock()
	defer c.loadMu.Unlock()

	c.mu.RLock()
	pages = append([]domain.SitePage(nil), c.pagesCache...)
	loadedAt = c.loadedAt
	c.mu.RUnlock()
	if len(pages) > 0 && time.Since(loadedAt) < c.cfg.SiteRefreshInterval {
		return pages, nil
	}

	freshPages, err := c.load(ctx)
	if err != nil {
		if len(pages) > 0 {
			return pages, nil
		}
		return nil, err
	}

	c.mu.Lock()
	c.pagesCache = append([]domain.SitePage(nil), freshPages...)
	c.loadedAt = time.Now()
	c.mu.Unlock()
	return append([]domain.SitePage(nil), freshPages...), nil
}

func (c *siteCatalog) load(ctx context.Context) ([]domain.SitePage, error) {
	localPages, localErr := content.LoadSitePages(c.cfg)
	if localErr == nil && len(localPages) > 0 {
		return localPages, nil
	}

	paths := c.cfg.SiteSeedPaths
	if len(paths) == 0 {
		paths = catalog.DefaultSeedPaths()
	}

	uniquePaths := uniqueNormalizedPaths(paths)
	if c.cfg.SiteMaxPages > 0 && len(uniquePaths) > c.cfg.SiteMaxPages {
		uniquePaths = uniquePaths[:c.cfg.SiteMaxPages]
	}
	if len(uniquePaths) == 0 {
		return nil, errors.New("no site seed paths configured")
	}

	type result struct {
		index int
		page  domain.SitePage
		err   error
	}

	workers := minInt(6, len(uniquePaths))
	jobs := make(chan int)
	results := make(chan result, len(uniquePaths))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for index := range jobs {
				page, err := c.fetchPage(ctx, uniquePaths[index])
				results <- result{index: index, page: page, err: err}
			}
		}()
	}

	for i := range uniquePaths {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	ordered := make(map[int]domain.SitePage, len(uniquePaths))
	for item := range results {
		if item.err != nil || strings.TrimSpace(item.page.Content) == "" {
			continue
		}
		ordered[item.index] = item.page
	}

	pages := make([]domain.SitePage, 0, len(ordered))
	for index := range uniquePaths {
		if page, ok := ordered[index]; ok {
			pages = append(pages, page)
		}
	}
	if len(pages) == 0 {
		if localErr != nil {
			return nil, fmt.Errorf("unable to load local content or fetch website pages: %w", localErr)
		}
		return nil, fmt.Errorf("unable to fetch any website pages from %s", c.cfg.SiteBaseURL)
	}
	return pages, nil
}

func (c *siteCatalog) fetchPage(ctx context.Context, path string) (domain.SitePage, error) {
	pageURL := strings.TrimRight(c.cfg.SiteBaseURL, "/") + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, pageURL, nil)
	if err != nil {
		return domain.SitePage{}, err
	}
	req.Header.Set("User-Agent", c.cfg.SiteUserAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	resp, err := c.client.Do(req)
	if err != nil {
		return domain.SitePage{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return domain.SitePage{}, fmt.Errorf("unexpected status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	if err != nil {
		return domain.SitePage{}, err
	}

	body := string(bodyBytes)
	title := extractTitle(body)
	content := extractVisibleText(body)
	content = truncateString(content, c.cfg.SiteMaxContentChars)
	if title == "" && content == "" {
		return domain.SitePage{}, errors.New("page did not contain readable content")
	}
	if title == "" {
		title = path
	}

	return domain.SitePage{
		URL:     pageURL,
		Path:    path,
		Title:   title,
		Summary: truncateString(content, 900),
		Content: content,
		Tokens:  tokenize(strings.Join([]string{title, path, content}, " ")),
	}, nil
}

func uniqueNormalizedPaths(paths []string) []string {
	seen := make(map[string]struct{}, len(paths))
	out := make([]string, 0, len(paths))
	for _, path := range paths {
		path = strings.TrimSpace(path)
		if path == "" {
			continue
		}
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		if _, ok := seen[path]; ok {
			continue
		}
		seen[path] = struct{}{}
		out = append(out, path)
	}
	return out
}

func extractTitle(body string) string {
	match := titleRe.FindStringSubmatch(body)
	if len(match) < 2 {
		return ""
	}
	return normalizeWhitespace(html.UnescapeString(match[1]))
}

func extractVisibleText(body string) string {
	body = commentRe.ReplaceAllString(body, " ")
	body = headRe.ReplaceAllString(body, " ")
	body = scriptRe.ReplaceAllString(body, " ")
	body = styleRe.ReplaceAllString(body, " ")
	body = noscriptRe.ReplaceAllString(body, " ")
	body = svgRe.ReplaceAllString(body, " ")
	body = tagRe.ReplaceAllString(body, " ")
	body = html.UnescapeString(body)
	return normalizeWhitespace(body)
}

func normalizeWhitespace(value string) string {
	value = whitespaceRe.ReplaceAllString(value, " ")
	return strings.TrimSpace(value)
}

func tokenize(value string) map[string]struct{} {
	value = strings.ToLower(value)
	matches := tokenRe.FindAllString(value, -1)
	if len(matches) == 0 {
		return nil
	}

	stopWords := map[string]struct{}{
		"the": {}, "and": {}, "for": {}, "that": {}, "with": {}, "this": {},
		"from": {}, "your": {}, "have": {}, "will": {}, "into": {}, "about": {},
		"what": {}, "when": {}, "where": {}, "which": {}, "their": {}, "they": {},
		"them": {}, "just": {}, "here": {}, "page": {}, "site": {},
	}

	out := make(map[string]struct{}, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		if _, skip := stopWords[match]; skip {
			continue
		}
		out[match] = struct{}{}
	}
	return out
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

func minInt(left, right int) int {
	if left < right {
		return left
	}
	return right
}
