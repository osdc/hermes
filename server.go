package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/osdc/hermes/api"
	"github.com/osdc/hermes/utils"
	"github.com/spf13/viper"
	"net/http"
    "fmt"
    "gopkg.in/redis.v5"

	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

)
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func AuthMiddleWare (next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            token := c.Request().Header.Get("Authorization")
            // TODO: Fix this crime
            if token == "" {
                return next(c)
            }
            userId, err := redisClient.Get(token).Result()

            if err == redis.Nil {
                response := make(map[string]interface{})
                response["status"] = "FAILED"
                response["error"] = "Auth Failed"
                return c.JSON(http.StatusUnauthorized, response)
            } else if err != nil {
                panic(err)
            }
            fmt.Println(userId)
            return next(c)
        }
    }


var redisClient = utils.GetRedisClient()

func main() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")


    err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read config file")
	}

	e := echo.New()

    e.Use(AuthMiddleWare)
    e.Use(ServerHeader)
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
