package models

import "time"

type Screencast struct {
	Title           string
	Date            time.Time
	Summary         string
	Complete        string
	YouTubeEmbedURL string
}
