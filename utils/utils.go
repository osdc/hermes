package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func GetDBConn() *gorm.DB {

	dbUser := viper.GetString("database.user")
	dbName := viper.GetString("database.name")
	dbPassword := viper.GetString("database.password")

	databaseCredentials := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", dbUser, dbName, dbPassword)
	db, err := gorm.Open("postgres", databaseCredentials)

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

func RequestWebkiosk(username, password string) (bool, string) {
	// password should be url encoded
	password = url.QueryEscape(password)

	webkioskURL := "https://webkiosk.jiit.ac.in/CommonFiles/UserAction.jsp?txtInst=Institute&InstCode=JIIT&txtuType=Member%%20Type%%20&UserType=S&txtCode=Enrollment%%20No&MemberCode=%s&txtPin=Password%%2FPin&Password=%s&BTNSubmit=Submit"

	reqURL := fmt.Sprintf(webkioskURL, username, password)

	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(respBody), "Invalid Password") {
		return false, "Invalid Password"
	}

	return true, "Valid Credentials"
}
