package actions

import "github.com/gobuffalo/buffalo"

// OtherHandler is a default handler to serve up
// a home page.
func OtherHandler(c buffalo.Context) error {
	name := c.Param("name")
	c.Set("name", name)
	return c.Render(200, r.HTML("other.html"))
}
