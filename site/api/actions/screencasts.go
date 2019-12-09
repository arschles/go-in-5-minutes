package actions

import (
	"log"

	"github.com/arschles/go-in-5-minutes/site/api/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
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

func screencastAddHook(c buffalo.Context) error {
return nil
}
