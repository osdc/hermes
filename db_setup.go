package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/osdc/hermes/models"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres dbname=hermes sslmode=disable password=qwe")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Hub{}, &models.Event{}, &models.Announcement{}, &models.Role{})
	db.Model(&models.User{}).AddForeignKey("User_ID", "Roles(User_Id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Hub{}).AddForeignKey("Hub_ID", "Roles(Hub_Id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Hub{}).AddForeignKey("Hub_ID", "Announcements(Hub_Id)", "RESTRICT", "RESTRICT")
}
