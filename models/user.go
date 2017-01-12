package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Slug         string `gorm:"primary_key"`
	Batch        string
	EnrollmentNo string `gorm:"unique"`
	About        string
	IsMaster     bool `gorm:"default:false"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Slug", "TODO")
	return nil
}
