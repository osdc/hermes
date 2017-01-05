package utils

import (
	"encoding/json"
	"io"
    "io/ioutil"
    "fmt"
    "net/http"
    "net/url"
    "log"

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

func RequestWebkiosk(username, dob, password string) bool {
    // dob: dd-mm-yyyy
    // password should be url encoded
    password = url.QueryEscape(password)

    webkioskURL := "https://webkiosk.jiit.ac.in/CommonFiles/UserActionn.jsp?x=&txtInst=Institute&InstCode=JIIT&txtuType=Member+Type&UserType=S&txtCode=Enrollment+No&MemberCode=%s&DOB=DOB&DATE1=%s&txtPin=Password%%2FPin&Password=%s&BTNSubmit=Submit"

    reqURL := fmt.Sprintf(webkioskURL, username, dob, password)

    resp, err := http.Get(reqURL)
    if err != nil {
        log.Fatal(err)
    }

    // TODO: @dwaipayan FIXME: Make cases for successful login, invalid dob
    //                         invalid password etc.
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(respBody)
    return true
}
