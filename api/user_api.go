package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/osdc/hermes/models"
	"github.com/osdc/hermes/utils"
)

func CreateUser(c echo.Context) error {

	payload := utils.ParseJSON(c.Request().Body)

	name := payload["name"]
	email := payload["email"]
	enroll := payload["enrollment_number"]
	password := payload["password"]
	batch := payload["batch"]

	db := utils.GetDBConn()
	defer db.Close()

    user := models.User{Name: name, Email: email,
		EnrollmentNo: enroll, Password: password,
		Batch: batch}

	db.Create(&user)

	response := utils.SuccessResponse()
	return c.JSON(http.StatusCreated, response)
}

func WebkioskAuth(c echo.Context) error {

	payload := utils.ParseJSON(c.Request().Body)

    enroll := payload["enrollment_number"]
    password := payload["password"]
    dob := payload["dob"]

    utils.RequestWebkiosk(enroll, dob, password)

    return c.JSON(http.StatusOK, utils.SuccessResponse())

}
