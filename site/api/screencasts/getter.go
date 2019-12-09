package screencasts

import (
	"time"

	"github.com/arschles/go-in-5-minutes/site/api/models"
)

func Get(env string) ([]*models.Screencast, error) {
	return []*models.Screencast{
		{
			Title:           "Test Screencast 1",
			Date:            time.Now(),
			Summary:         "This is a summary of the test screencast",
			Complete:        "This is the complete text of the test screencast",
			YouTubeEmbedURL: "https://goin5minutes.com",
		},
	}, nil
}
