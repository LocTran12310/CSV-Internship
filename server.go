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
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24}) // Expire in a day
	router.Use(sessions.Sessions("mysession", store))

	// USERS API
	router.POST("/login", repos.Login)
	router.GET("/logout", repos.Logout)
	router.POST("/user/signup", repos.CreateUser)

	// Authentication
	auth := router.Group("auth")
	auth.Use(middleware.Authentication())
	auth.PATCH("/user/changepassword/:id", repos.ChangePassword)
	auth.Static("/user", "public/users")

	// PROFILES API
	profiles := auth.Group("profiles")
	profiles.GET("/", repos.GetProfiles)
	profiles.GET("/:id", repos.GetProfile)
	profiles.POST("/create", repos.CreateProfile)
	profiles.PATCH("/update/:id", repos.UpdateProfile)
	profiles.PATCH("/delete/:id", repos.DeleteProfile)

	// COURSE API
	courses := auth.Group("courses")
	courses.GET("/", repos.GetCourses)
	courses.GET("/:id", repos.GetCourse)
	courses.GET("/:id/participants", repos.GetCourseParticipants)
	courses.GET("/participants", repos.GetAllCourseParticipants)
	courses.POST("/create", repos.CreateCourse)
	courses.PATCH("/update/:id", repos.UpdateCourse)
	courses.PATCH("/delete/:id", repos.DeleteCourse)
	courses.PATCH("/participant/delete/:id", repos.DeleteCourseParticipant)
	courses.POST("/participant/create", repos.CreateCourseParticipant)

	//POSITION API
	positions := auth.Group("positions")
	positions.GET("/", repos.GetPositions)
	positions.POST("/create", repos.CreatePosition)
	positions.PATCH("/update/:id", repos.UpdatePosition)
	positions.PATCH("/delete/:id", repos.DeletePosition)

	//DEPARTMENT API
	departments := auth.Group("departments")
	departments.GET("/", repos.GetDepartments)
	departments.POST("/create", repos.CreateDepartment)
	departments.PATCH("/update/:id", repos.UpdateDepartment)
	departments.PATCH("/delete/:id", repos.DeleteDepartment)

	// LEAVE API
	leave := auth.Group("leave")
	leave.GET("/", repos.GetAllLeaveDetails)
	leave.GET("/:employee_id", repos.GetLeaveDetails)
	leave.POST("/create", repos.CreateLeaveDetails)
	leave.PATCH("/update/:id", repos.UpdateLeaveDetails)
	leave.PATCH("/delete/:id", repos.DeleteLeaveDetails)

	leaveTypes := leave.Group("types")
	leaveTypes.GET("/", repos.GetLeaveTypes)
	leaveTypes.POST("/create", repos.CreateLeaveType)
	leaveTypes.PATCH("/update/:id", repos.UpdateLeaveType)
	leaveTypes.PATCH("/delete/:id", repos.DeleteLeaveType)

	leaveReasons := leave.Group("reasons")
	leaveReasons.GET("/", repos.GetLeaveReasons)
	leaveReasons.POST("/create", repos.CreateLeaveReason)
	leaveReasons.PATCH("/update/:id", repos.UpdateLeaveReason)
	leaveReasons.PATCH("/delete/:id", repos.DeleteLeaveReason)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
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
