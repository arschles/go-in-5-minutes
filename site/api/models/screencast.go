package models

import "time"

type Screencast struct {
	Title           string
	Data            time.Time
	Summary         string
	Complete        string
	YouTubeEmbedURL string
}
