package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.AddTrailingSlash())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} | ${uri} | ${status} | ${error}\n",
	}))
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	app.GET("/", func(c echo.Context) error {
		return c.JSON(200, "HIT")
	})
	log.Fatal(app.Start(":3000"))
}
