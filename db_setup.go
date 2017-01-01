package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wimpykid26/hermes/models"
)

func main() {
	db, err := gorm.Open("postgres", "user=<yoursudouser> dbname=hermes sslmode=disable password=<yoursudopassword>")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&models.Users{}, &models.Hubs{}, &models.Events{}, &models.Announcements{}, &models.Roles{})
	db.Model(&models.Users{}).AddForeignKey("User_ID", "Roles(User_Id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Hubs{}).AddForeignKey("Hub_ID", "Roles(Hub_Id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Hubs{}).AddForeignKey("Hub_ID", "Announcements(Hub_Id)", "RESTRICT", "RESTRICT")
}
