package main

import (
	"github.com/labstack/echo"
	"github.com/osdc/hermes/api"
    "github.com/spf13/viper"
	"net/http"
)

func main() {
    viper.SetConfigName("app")
    viper.AddConfigPath("config")

    err := viper.ReadInConfig()
    if err != nil {
        panic("cannot read config file")
    }

	e := echo.New()
	e.POST("/api/user/new", api.CreateUser)
	e.POST("/api/user/webkioskauth", api.WebkioskAuth)
	//e.GET("/insert_hub", api.Insert_hub)
	//e.GET("/insert_event", api.Insert_event)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":4000")
}
