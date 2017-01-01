package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Hub struct {
	gorm.Model
	Name  string
	About string
	Slug  string `gorm:"primary_key"`
}

type Event struct {
	gorm.Model
	Name        string
	HubID       uint // TODO Add Foreign Key
	Event       string
	Description string
	Date        time.Time
}

type Announcement struct {
	gorm.Model
	HubID       uint `gorm:"not null;unique"`
	Description string
}

type Role struct {
	gorm.Model
	UserID uint `gorm:"not null;unique"`
	HubID  uint `gorm:"not null;unique"`
	Role   string
}
