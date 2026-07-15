package readcount

import "time"

type ReadCountLog struct {
	Id           int64
	ItemId       string
	Identity     string
	VisitorID    string
	VisitorIp    string
	UserAgent    string
	Referer      string
	ReadDuration int
	ReadDepth    int
	CreatedAt    time.Time
}

type ReadCountRequest struct {
	ItemID       string `json:"itemId"`
	ItemType     int    `json:"itemType"`
	Identity     string
	VisitorID    string `json:"visitorId"`
	UserID       string `json:"userId"`
	IP           string
	UserAgent    string
	Referer      string
	ReadDuration int `json:"readDuration"`
	ScrollDepth  int `json:"scrollDepth"`
}
