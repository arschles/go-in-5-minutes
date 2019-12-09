package models

import "time"

type Screencast struct {
	Title           string
	Date            time.Time
	Summary         string
	Complete        string
	YouTubeEmbedURL string
}

func (s Screencast) HumanTime() string {
	return s.Date.Format("Mon Jan 2, 2006")
}
