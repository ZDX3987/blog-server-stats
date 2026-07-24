package pageview

import "time"

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
	OccurredAt     time.Time
}

type PageViewRequest struct {
	EventID        string `json:"eventId"`
	VisitorID      string
	Path           string `json:"path"`
	Referer        string
	UserAgent      string
	IP             string
	DurationSec    int       `json:"durationSec"`
	MaxScrollDepth uint8     `json:"maxScrollDepth"`
	OccurredAt     time.Time `json:"occurredAt"`
}

type PageViewMultiRequest struct {
	VisitorID string
	Referer   string
	UserAgent string
	IP        string
	Request   []PageViewRequest
}
