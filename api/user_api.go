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
	enroll := payload["enrollment_number"]
	password := payload["password"]
	batch := payload["batch"]

	db := utils.GetDBConn()
	defer db.Close()

    user := models.User{Name: name,
		EnrollmentNo: enroll, Password: password,
		Batch: batch}

    authentic, msg := utils.RequestWebkiosk(enroll, dob, password)

    if !authentic {
        response := make(map[string]string)
        response["status"] = "Failed"
        response["error"] = msg
        return c.JSON(http.StatusUnauthorized, response)
    }

    db.Create(&user)

	response := utils.SuccessResponse()
	return c.JSON(http.StatusCreated, response)
}

func WebkioskAuth(c echo.Context) error {

	payload := utils.ParseJSON(c.Request().Body)

    enroll := payload["enrollment_number"]
    password := payload["password"]
    dob := payload["dob"]

    authentic := utils.RequestWebkiosk(enroll, dob, password)

    response := make(map[string]string)
    response["status"] = "OK"
    response["auth"] = authentic

    return c.JSON(http.StatusOK, response)
}
