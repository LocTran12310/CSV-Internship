package main

import (
	"golangapi/middleware"
	repos "golangapi/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// USERS API
	router.POST("/login", repos.Login)
	router.GET("/logout", repos.Logout)
	router.POST("/user/signup", repos.CreateUser)

	// Authentication
	auth := router.Group("/profiles")
	auth.Use(middleware.Authentication())

	// PROFILES API
	auth.GET("/", repos.GetProfiles)
	auth.GET("/:id", repos.GetProfile)
	auth.POST("/create", repos.CreateProfile)
	auth.PATCH("/update/:id", repos.UpdateProfile)
	auth.PATCH("/delete/:id", repos.DeleteProfile)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
