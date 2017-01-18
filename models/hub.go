package models

import "github.com/jinzhu/gorm"

type Hub struct {
	gorm.Model
	Name  string
	About string
	Slug  string `gorm:"not null;unique"`
}

func (hub Hub) SerializeHub() map[string]interface{} {
	hubData := make(map[string]interface{})

	hubData["name"] = hub.Name
	hubData["id"] = hub.ID
	hubData["about"] = hub.About
	hubData["slug"] = hub.Slug

	return hubData
}
