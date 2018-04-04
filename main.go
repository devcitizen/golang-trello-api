package main

import (
	"fmt"
	ctrl "savannah-go/controllers"
	m "savannah-go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:12345678@/savannah-go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	var user m.User
	var role m.Role
	var sprint m.Sprint
	var project m.Project
	var backlog m.Backlog

	db.AutoMigrate(&user, &role, &sprint, &project, &backlog)
}

func main() {
	router := gin.Default()
	router.Use(func(context *gin.Context) {
		// add header Access-Control-Allow-Origin
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
		context.Next()
	})

	users := router.Group("/api/v1/users")
	{
		users.GET("/", ctrl.GetUsers)
		users.POST("/", ctrl.CreateUser)
		users.GET("/:id", ctrl.GetUser)
		users.PUT("/:id", ctrl.UpdateUser)
		users.DELETE("/:id", ctrl.DeleteUser)
	}

	roles := router.Group("api/v1/roles")
	{
		roles.GET("/", ctrl.GetRoles)
		roles.GET("/:id", ctrl.GetRole)
		roles.POST("/", ctrl.PostRole)
		roles.DELETE("/:id", ctrl.DeleteRole)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
