package model

import "time"

type ReadCountRequest struct {
	ReadDuration int `json:"readDuration"`
	ReadDepth    int `json:"readDepth"`
}

type ReadCountLog struct {
	Id           int64
	ItemId       string
	Identity     string
	VisitorId    string
	VisitorIp    string
	UserAgent    string
	Referer      string
	ReadDuration int
	ReadDepth    int
	CreatedAt    time.Time
}

type ReadCountForm struct {
	ItemId string `json:"itemId"`
}
