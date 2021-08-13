package main

import (
	"golangapi/apis/profileapi"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/profiles", profileapi.GetAllProfiles)
	router.POST("/profiles/create", profileapi.AddProfile)
	router.PATCH("/profiles/update/:id", profileapi.UpdateProfile)

	router.Run("localhost:8888")
}
