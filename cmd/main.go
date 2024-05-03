package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mccune1224/flashkick/pkg/start"
)

func main() {
	app := echo.New()
	app.Use(middleware.AddTrailingSlash())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} | ${uri} | ${status} | ${error}\n",
	}))
	app.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	startGraphQL := start.NewStartGGClient(os.Getenv("START_GG_TOKEN"))

	app.GET("/", func(c echo.Context) error {
		var MeQuery struct {
			Me struct {
				Name string
			}
		}
		err := startGraphQL.Client.Query(c.Request().Context(), MeQuery, nil)
		if err != nil {
			c.JSON(500, err.Error())
		}
		log.Println(MeQuery)
		return c.JSON(200, "HIT")
	})

	log.Fatal(app.Start(":3000"))
}
