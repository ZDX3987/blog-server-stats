package pageview

type PageView struct {
	ID             int
	EventID        string
	VisitorID      string
	Path           string
	Referer        string
	UserAgent      string
	IP             string
	DurationSec    int
	MaxScrollDepth uint8
}

type PageViewRequest struct {
	EventID        string `json:"event_id"`
	VisitorID      string
	Path           string `json:"path"`
	Referer        string
	UserAgent      string
	IP             string
	DurationSec    int   `json:"duration_sec"`
	MaxScrollDepth uint8 `json:"max_scroll_depth"`
}
