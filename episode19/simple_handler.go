package main

import (
	"net/http"
	"strconv"

	"github.com/gedex/inflector"
	"github.com/labstack/echo"
)

func pluralizeHandler(c echo.Context) error {
	singular := c.QueryParam("singular")
	numStr := c.QueryParam("num")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		c.Logger().Errorf("'num' query param %s is not a number", numStr)
		return err
	}

	// singular case
	if num == 1 {
		return c.String(http.StatusOK, singular)
	}
	// plural case
	return c.String(http.StatusOK, inflector.Pluralize(singular))
}
