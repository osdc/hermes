package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/spf13/viper"

	"github.com/osdc/hermes/models"
)

func main() {
    viper.SetConfigName("app")
    viper.AddConfigPath("config")

    err := viper.ReadInConfig()
    if err != nil {
        panic("Config file not found")
    }

    dbUser := viper.GetString("database.user")
    dbName := viper.GetString("database.name")
    dbPassword := viper.GetString("database.password")

    databaseCredentials := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbUser, dbName, dbPassword)
		db, err := gorm.Open("postgres", databaseCredentials)

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
