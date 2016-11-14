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
	reqCounter := requestCounter{}
	e.Use(reqCounter.Process)

	/////
	// path and query params (simple)
	/////
	e.GET(fmt.Sprintf("/pluralize/:%s", singularPathParam), pluralizeHandler)

	/////
	// number of incoming requests, using the request counter middleware
	/////
	e.GET("/request_count", reqCounter.handle)

	/////
	// accept a JSON body
	/////
	e.POST("/json", jsonHandler)

	const port = 8080
	e.Logger.Printf("starting on port %d", port)
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
