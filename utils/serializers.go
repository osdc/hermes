package utils

import (
  "github.com/osdc/hermes/models"
)

func SerializeHub(hub models.Hub) map[string]interface{} {
  hubData := make(map[string]interface{})

	hubData["name"] = hub.Name
	hubData["id"] = hub.ID
	hubData["about"] = hub.About
	hubData["slug"] = hub.Slug

  return hubData;
}

func SerializeUser(user models.User) map[string]interface{} {
  userData := make(map[string]interface{})

	userData["name"] = user.Name
	userData["id"] = user.ID
	userData["email"] = user.Email
	userData["batch"] = user.Batch

  return userData;
}
