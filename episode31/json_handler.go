package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type coordinates struct {
	X int `json:"x" xml:"x" form:"x"`
	Y int `json:"y" xml:"y" form:"y"`
}

func jsonHandler(c echo.Context) error {
	coords := new(coordinates)
	if err := c.Bind(coords); err != nil {
		return err
	}
	return c.String(
		http.StatusOK,
		fmt.Sprintf("x coord = %d, y coord = %d", coords.X, coords.Y),
	)
}
