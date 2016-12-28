package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type Users struct {
	gorm.Model
	User_Id      int    `gorm:"AUTO_INCREMENT"`
	Name         string `gorm:"primary_key"`
	Batch        string
	is_master    string
	EnrollmentNo string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Hubs struct {
	gorm.Model
	Hub_Id    int    `gorm:"AUTO_INCREMENT"`
	Name      string `gorm:"primary_key"`
	Bio       string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Events struct {
	gorm.Model
	Event_Id    int    `gorm:"AUTO_INCREMENT"`
	Name        string `gorm:"primary_key"`
	Event       string
	Description string
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Announcements struct {
	gorm.Model
	hub_id      int
	description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Roles struct {
	gorm.Model
	user_id   int
	hub_id    int
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func insert_user(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "hermes.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	name := c.QueryParam("name")
	batch := c.QueryParam("batch")
	is_master := c.QueryParam("is_master")
	bio := c.QueryParam("bio")
	enrollmentNo := c.QueryParam("enrollmentNo")
	db.Create(&Users{Name: name, Batch: batch, is_master: is_master, EnrollmentNo: enrollmentNo, Bio: bio, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	return c.String(http.StatusOK, "Inserted User Successfully")
}

func insert_hub(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "hermes.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	name := c.QueryParam("name")
	bio := c.QueryParam("bio")
	db.Create(&Hubs{Name: name, Bio: bio, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	return c.String(http.StatusOK, "Inserted Hub Successfully")
}

func insert_event(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "hermes.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	name := c.QueryParam("name")
	event := c.QueryParam("event")
	description := c.QueryParam("description")
	date := time.Now()
	db.Create(&Events{Name: name, Event: event, Description: description, Date: date, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	return c.String(http.StatusOK, "Inserted Event Successfully")
}

func main() {
	e := echo.New()
	e.GET("/insert_user", insert_user)
	e.GET("/insert_hub", insert_hub)
	e.GET("/insert_event", insert_event)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
