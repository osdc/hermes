package utils

import (
	"encoding/json"
	"io"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDBConn() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres dbname=hermes sslmode=disable password=qwe")

	if err != nil {
		panic("[ERROR] Database connection failed")
	}

	return db
}

func ParseJSON(form io.Reader) map[string]string {

	data := make(map[string]string)
	err := json.NewDecoder(form).Decode(&data)

	if err != nil {
		// TODO
		panic("Improve Error Message")
	}

	return data
}

func SuccessResponse() map[string]string {
	response := make(map[string]string)
	response["status"] = "OK"
	return response
}
