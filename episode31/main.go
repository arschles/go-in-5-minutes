package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	/////
	// Setting up basic middleware(s). There are lots of other built-in middlewares, including:
	//
	// 1. CSRF prevention: https://echo.labstack.com/middleware/csrf
	// 2. CORS support: https://echo.labstack.com/middleware/cors
	// 3. JWT support: https://echo.labstack.com/middleware/jwt
	// 4. HTTPS redirects: https://echo.labstack.com/middleware/redirect
	/////
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	reqCounter := requestCounter{}
	e.Use(reqCounter.Process)

	/////
	// path and query params (simple)
	/////
	e.GET(fmt.Sprintf("/pluralize/:%s", singularPathParam), pluralizeHandler)

	/////
	// use the request counter middleware to return the number of incoming requests
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
