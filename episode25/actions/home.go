package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// HomePostHandler is a handler to serve the POST / page
func HomePostHandler(c buffalo.Context) error {
	var name string
	if c.Param("name") != "" {
		name = c.Param("name")
		c.Set("name", name)
		return c.Render(200, r.HTML("index.html"))
	}
	return c.Redirect(http.StatusFound, "/")
}
