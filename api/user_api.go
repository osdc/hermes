package api

import (
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	"github.com/osdc/hermes/models"
	"github.com/osdc/hermes/utils"
)

func CreateUser(c echo.Context) error {
	payload := utils.ParseJSON(c.Request().Body)

	name := payload["name"]
	enroll := payload["enrollment_number"]
	password := payload["password"]
	batch := payload["batch"]
	dob := payload["dob"]

	password = []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	db := utils.GetDBConn()
	defer db.Close()

	user := models.User{Name: name, EnrollmentNo: enroll, Password: string(hashedPassword), Batch: batch}

	authentic, msg := utils.RequestWebkiosk(enroll, dob, password)

	if !authentic {
		response := make(map[string]interface{})
		response["status"] = "FAILED"
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

	authentic, msg := utils.RequestWebkiosk(enroll, dob, password)

	if !authentic {
		response := make(map[string]interface{})
		response["status"] = "FAILED"
		response["error"] = msg
		return c.JSON(http.StatusUnauthorized, response)
	}

	response := make(map[string]interface{})
	response["status"] = "OK"
	response["auth"] = authentic

	return c.JSON(http.StatusOK, response)
}

func LoginUser(c echo.Context) error {
	var user models.User

	payload := utils.ParseJSON(c.Request().Body)

	enroll := payload["enrollment_number"]
	password := payload["password"]

	db := utils.GetDBConn()
	defer db.Close()

	db.Where("enrollment_no = ?", enroll, password).First(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))

	if user.ID == 0 && err != nil {
		response := make(map[string]interface{})
		response["status"] = "FAILED"
		response["error"] = "Wrong Credentials"
		return c.JSON(http.StatusUnauthorized, response)
	}

	response := make(map[string]interface{})
	// TODO: Write Serializers
	userData := make(map[string]interface{})
	userData["name"] = user.Name
	userData["id"] = user.ID
	userData["email"] = user.Email
	userData["batch"] = user.Batch
	response["status"] = "OK"
	response["user"] = userData

	return c.JSON(http.StatusOK, response)
}
