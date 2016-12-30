package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/wimpykid26/hermes/models"
	"net/http"
)

func Insert_user(c echo.Context) error {
	db, err := gorm.Open("postgres", "user=<yoursudouser> dbname=hermes sslmode=disable password=<yoursudopassword>")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	name := c.QueryParam("name")
	batch := c.QueryParam("batch")
	is_master := c.QueryParam("is_master")
	bio := c.QueryParam("bio")
	enrollmentNo := c.QueryParam("enrollmentNo")
	db.Create(&models.Users{Name: name, Batch: batch, Is_master: is_master, EnrollmentNo: enrollmentNo, Bio: bio})
	return c.String(http.StatusOK, "Inserted User Successfully")
}

func Insert_hub(c echo.Context) error {
	db, err := gorm.Open("postgres", "user=<yoursudouser> dbname=hermes sslmode=disable password=<yoursudopassword>")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	name := c.QueryParam("name")
	bio := c.QueryParam("bio")
	db.Create(&models.Hubs{Name: name, Bio: bio})
	return c.String(http.StatusOK, "Inserted Hub Successfully")
}

func Insert_event(c echo.Context) error {
	db, err := gorm.Open("postgres", "user=<yoursudouser> dbname=hermes sslmode=disable password=<yoursudopassword>")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	name := c.QueryParam("name")
	event := c.QueryParam("event")
	date := c.QueryParam("date")
	description := c.QueryParam("description")
	db.Create(&models.Events{Name: name, Event: event, Description: description, Date: date})
	return c.String(http.StatusOK, "Inserted Event Successfully")
}
