package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Users struct {
	gorm.Model
	User_ID      int    `gorm:"AUTO_INCREMENT"`
	Name         string `gorm:"primary_key"`
	Batch        string
	Is_master    string
	EnrollmentNo string
	Bio          string
}

type Hubs struct {
	gorm.Model
	Hub_ID int    `gorm:"AUTO_INCREMENT"`
	Name   string `gorm:"primary_key"`
	Bio    string
}

type Events struct {
	gorm.Model
	Event_ID    int    `gorm:"AUTO_INCREMENT"`
	Name        string `gorm:"primary_key"`
	Event       string
	Description string
	Date        string
}

type Announcements struct {
	gorm.Model
	HubId       int `gorm:"not null;unique"`
	Description string
}

type Roles struct {
	gorm.Model
	UserId int `gorm:"not null;unique"`
	HubId  int `gorm:"not null;unique"`
	Role   string
}
