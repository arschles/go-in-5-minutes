package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo"
)

// requestCounter is both a middleware that counts the number of incoming requests, and a handler that can return that number
type requestCounter struct {
	counter int64
}

func (r *requestCounter) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		atomic.AddInt64(&r.counter, 1)
		return next(c)
	}
}

func (r *requestCounter) handle(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%d", atomic.LoadInt64(&r.counter)))
}
