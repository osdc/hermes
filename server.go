package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/osdc/hermes/api"
)

func main() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read config file")
	}

	e := echo.New()

	e.Use(middleware.CORS())

	// User Actions
	e.POST("/api/user/new", api.CreateUser)
	// e.POST("/api/user/webkioskauth", api.WebkioskAuth)
	e.POST("/api/user/login", api.LoginUser)

	// Hub Actions
	e.POST("/api/hub/new", api.CreateHub)
	e.GET("/api/hub/:slug", api.GetHub)
	e.GET("/api/hub", api.GetAllHubs)

	// Miscellaneous Actions
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":4000")
}
