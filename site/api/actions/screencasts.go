package actions

import (
	"github.com/arschles/go-in-5-minutes/site/api/screencasts"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
)

func screencastSummaryListHandler(c buffalo.Context) error {
	casts, err := screencasts.Get(envy.Get("GO_ENV", "development"))
	if err != nil {
		return err
	}
	c.Set("screencasts", casts)
	return c.Render(200, r.HTML("screencast_summary_list.html"))
}
