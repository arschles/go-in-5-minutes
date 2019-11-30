package screencasts

func Get(env string) ([]*models.Screencast, error) {
	return []&models.Screencast{
		Title:"Test Screencast 1",
		Data: time.Now(),
		Summary: "This is a summary of the test screencast",
		Complete: "This is the complete text of the test screencast",
		YoutubeEmbedURL: "https://goin5minutes.com",
	}
}
