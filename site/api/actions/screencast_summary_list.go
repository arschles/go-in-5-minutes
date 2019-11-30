package actions

import (
	"github.com/gobuffalo/buffalo"
)

func screencastSummaryListHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("screencast_summary_list.html"))
}
