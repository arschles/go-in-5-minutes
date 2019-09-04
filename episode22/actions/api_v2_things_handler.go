package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func apiV2ThingsHandler(c buffalo.Context) error {
	things := []string{"thing4", "thing5", "thing6"}
	return c.Render(http.StatusOK, r.JSON(things))
}
