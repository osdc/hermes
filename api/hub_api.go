package api

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/osdc/hermes/models"
	"github.com/osdc/hermes/utils"
)

func CreateHub(c echo.Context) error {
	payload := utils.ParseJSON(c.Request().Body)

	name := payload["name"]
	about := payload["about"]
	slug := payload["slug"]

	db := utils.GetDBConn()
	defer db.Close()

	hub := models.Hub{Name: name, About: about, Slug: slug}

	err := db.Create(&hub)

	if err.Error != nil {
		response := make(map[string]interface{})
		response["status"] = "FAILED"
		response["error"] = "Hub with similar slug already exists"
		return c.JSON(http.StatusConflict, response)
	}

	response := utils.SuccessResponse()
	return c.JSON(http.StatusCreated, response)
}

func GetHub(c echo.Context) error {
	var hub models.Hub

	slug := c.Param("slug")

	db := utils.GetDBConn()
	defer db.Close()

	db.Where("slug = ?", slug).First(&hub)

	if hub.ID == 0 {
		response := make(map[string]interface{})
		response["status"] = "FAILED"
		response["error"] = "Hub not found"
		return c.JSON(http.StatusNotFound, response)
	}

	response := make(map[string]interface{})
	response["status"] = "OK"
	response["hub"] = hub.Serialize()

	return c.JSON(http.StatusOK, response)
}

func GetAllHubs(c echo.Context) error {
	var hubs []models.Hub

	db := utils.GetDBConn()
	defer db.Close()

	db.Find(&hubs)

	hubsData := make([]interface{}, len(hubs))

	for i := 0; i < len(hubs); i++ {
		hubsData[i] = hubs[i].Serialize()
	}

	response := make(map[string]interface{})
	response["status"] = "OK"
	response["hubs"] = hubsData

	return c.JSON(http.StatusOK, response)
}
