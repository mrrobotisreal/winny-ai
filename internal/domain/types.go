package domain

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages  []Message `json:"messages"`
	PageURL   string    `json:"pageUrl,omitempty"`
	PageTitle string    `json:"pageTitle,omitempty"`
	PagePath  string    `json:"pagePath,omitempty"`
}

type StreamEvent struct {
	Type      string `json:"type"`
	Assistant string `json:"assistant,omitempty"`
	Model     string `json:"model,omitempty"`
	Delta     string `json:"delta,omitempty"`
	Error     string `json:"error,omitempty"`
}

type ChatSettings struct {
	TextModel     string
	Temperature   float64
	TopP          float64
	TopK          int
	NumCtx        int
	NumPredict    int
	RepeatPenalty float64
	KeepAlive     string
	SystemPrompt  string
}

type AssistantDelta struct {
	Content string
	Model   string
	Done    bool
}

type AssistantResponse struct {
	Content string
	Model   string
}

type SitePage struct {
	URL     string
	Path    string
	Title   string
	Summary string
	Content string
	Tokens  map[string]struct{} `json:"-"`
}
