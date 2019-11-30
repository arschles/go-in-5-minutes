package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/arschles/go-in-5-minutes/site/api/screencasts"
)

func screencastSummaryListHandler(c buffalo.Context) error {
	casts, err := screencasts.Get(envy.Get("GO_ENV"))
	return c.Render(200, r.HTML("screencast_summary_list.html"))
}
