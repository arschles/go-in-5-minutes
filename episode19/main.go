package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	/////
	// Setting up basic middleware(s)
	/////
	e.Use(middleware.Logger())

	/////
	// simple example
	/////
	e.GET(fmt.Sprintf("/pluralize/:%s", singularPathParam), pluralizeHandler)

	const port = 8080
	e.Logger.Printf("starting on port %d", port)
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
