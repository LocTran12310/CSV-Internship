package main

import (
	repos "golangapi/repository"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// TODO: add routes to router
	router.GET("/profiles", repos.GetProfiles)
	router.GET("/profiles/:id", repos.GetProfile)
	router.POST("/profiles/create", repos.CreateProfile)
	router.PATCH("/profiles/update/:id", repos.UpdateProfile)
	router.PATCH("/profiles/delete/:id", repos.DeleteProfile)

	return router
}
