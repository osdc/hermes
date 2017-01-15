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

func ShowHub(c echo.Context) error {
	var hub models.Hub

	payload := utils.ParseJSON(c.Request().Body)

	slug := payload["slug"]

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
	// TODO: Write Serializers
	hubData := make(map[string]interface{})
	hubData["name"] = hub.Name
	hubData["id"] = hub.ID
	hubData["about"] = hub.About
	hubData["slug"] = hub.Slug
	response["status"] = "OK"
	response["hub"] = hubData

	return c.JSON(http.StatusOK, response)
}

func ShowHubs(c echo.Context) error {
	var hubs []models.Hub

	db := utils.GetDBConn()
	defer db.Close()

	db.Find(&hubs)

	hubsData := make([]interface{}, len(hubs))

	for i := 0; i < len(hubs); i++ {
		hubData := make(map[string]interface{})

		hubData["name"] = hubs[i].Name
		hubData["id"] = hubs[i].ID
		hubData["about"] = hubs[i].About
		hubData["slug"] = hubs[i].Slug

		hubsData[i] = hubData
	}

	response := make(map[string]interface{})
	// TODO: Write Serializers
	response["status"] = "OK"
	response["hubs"] = hubsData

	return c.JSON(http.StatusOK, response)
}
