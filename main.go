package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.SetHTTPErrorHandler(func(err error, ctx echo.Context) {
		e.DefaultHTTPErrorHandler(err, ctx)
	})

	e.Get("/", accessible)

	e = hs256(e)
	e = rs256(e)

	e.Run(standard.New(":1323"))
}
