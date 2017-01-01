package main

import (
	"github.com/labstack/echo"
	"github.com/osdc/hermes/api"
	"net/http"
)

func main() {
	e := echo.New()
	e.POST("/insert_user", api.CreateUser)
	//e.GET("/insert_hub", api.Insert_hub)
	//e.GET("/insert_event", api.Insert_event)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":4000")
}
