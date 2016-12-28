package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	Hub_Id int    `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"primary_key"`
	Bio    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Events struct {
	gorm.Model
	Event_Id    int    `gorm:"AUTO_INCREMENT"`
	Name        string `gorm:"primary_key"`
	Event       string
	Description string
	Date        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Announcements struct {
	gorm.Model
	hub_id      int
	description string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Roles struct {
	gorm.Model
	user_id int
	hub_id  int
	Role    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	db, err := gorm.Open("sqlite3", "hermes.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Users{}, &Hubs{}, &Announcements{}, &Events{}, &Roles{})
}
