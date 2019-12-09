package actions

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/arschles/go-in-5-minutes/site/api/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/google/go-github/github"
)

func screencastSummaryListHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	casts := []models.Screencast{}
	q := tx.Q()
	q.Order("date desc")
	if err := q.All(&casts); err != nil {
		log.Printf("(screencastSummaryHandler) Error getting all screencasts\n%s", err)
		return err
	}

	c.Set("screencasts", casts)
	return c.Render(200, r.HTML("screencast_summary_list.html"))
}

func getScreencast(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	id := c.Param("id")
	cast := &models.Screencast{}
	if err := tx.Find(cast, id); err != nil {
		log.Printf("(getScreencast) Error getting screencast %s\n%s", id)
		return err
	}
	c.Set("screencast", cast)
	return c.Render(200, r.HTML("screencast_full.html"))
}

func screencastAddHook(ghClient *github.Client) func(buffalo.Context) error {
	return func(c buffalo.Context) error {
		tx := c.Value("tx").(*pop.Connection)

		cast := &models.Screencast{
			Date: time.Now(),
		}
		episodeID := c.Param("episode_id")
		if episodeID == "" {
			return c.Error(400, errors.New("episode number is required"))
		}
		readmePath := fmt.Sprintf("episode%s/README.md", episodeID)
		readmeMD, err := getFileFromGH(c, ghClient, readmePath)
		if err != nil {
			log.Printf("(screencastAddHook) Error getting episode %s readme from GH\n%s", episodeID, err)
			return err
		}

		ytFilePath := fmt.Sprintf("episode%s/youtube_url.txt", episodeID)
		ytFile, err := getFileFromGH(c, ghClient, ytFilePath)
		if err != nil {
			log.Printf("(screencastAddHook) Error getting episode %s youtube file from GH\n%s", episodeID, err)
			return err
		}

		html, _, err := ghClient.Markdown(c, readmeMD, &github.MarkdownOptions{
			Mode:    "gfm",
			Context: "google/go-github",
		})
		if err != nil {
			log.Printf("(getScreencast) error parsing README markdown for episode %s\n%s", episodeID, err)
			return err
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			log.Printf("(getScreencast) error parsing HTML for episode %s\n%s", episodeID, err)
			return err
		}

		// title
		title := doc.Find("h1").Text()
		if title == "" {
			log.Printf("(getScreencast) couldn't find title for episode %s\n%s", episodeID, err)
			return err
		}
		cast.Title = title

		// summary
		summary := doc.Find("p").First().Text()[:100]
		cast.Summary = summary

		// complete
		cast.Complete = html

		cast.YouTubeEmbedURL = ytFile

		verrs, err := tx.ValidateAndCreate(cast)
		if err != nil {
			log.Printf("(getScreencast) couldn't validate and create episode %s in the DB\n%s\n%s", episodeID, verrs, err)
			return err
		}
		return nil
	}

}
