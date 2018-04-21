package main

import (
	"fmt"
	"net/http"
	"os"
	ctrl "savannah-go/controllers"
	m "savannah-go/models"
	"time"
	u "zaraaya/utils"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	MysqlURL = "root:12345678@/savannah-go?charset=utf8&parseTime=True&loc=Local"
)

var db *gorm.DB
var user m.User
var role m.Role
var sprint m.Sprint
var project m.Project
var backlog m.Backlog

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:3002"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:  []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "ID-Company", "ID-Store", "ID-Merchant"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	r.GET("/ping", ping)
	//open a db connection
	var err error
	uri := os.Getenv("MYSQL_URL")

	if len(uri) == 0 {
		uri = MysqlURL
	}

	db, err := gorm.Open("mysql", uri)
	if err != nil {
		panic("failed to connect database")
	}
	// the jwt middleware

	v1 := r.Group("/api/v1")
	v1.Use(ConnectMiddleware(db))
	{
		authMiddleware := &jwt.GinJWTMiddleware{
			Realm:      "SAVANNAH",
			Key:        []byte("secret key"),
			Timeout:    time.Hour,
			MaxRefresh: time.Hour,
			Authenticator: func(email string, password string, c *gin.Context) (string, bool) {

				if err := db.Where("email = ?", email).First(&user).Error; err == nil {
					fmt.Println(email)
					match := u.CheckPasswordHash(password, user.Password)
					if match {
						return email, true
					}
				}

				return email, false
			},
			Unauthorized: func(c *gin.Context, code int, message string) {
				c.JSON(code, gin.H{
					"code":    code,
					"message": message,
				})
			},
			TokenLookup:   "header:Authorization",
			TokenHeadName: "Bearer",
			TimeFunc:      time.Now,
		}
		db.AutoMigrate(&user, &role, &sprint, &project, &backlog)

		r.POST("/login", authMiddleware.LoginHandler)
		//grouping route for api todos
		users := v1.Group("/users")
		users.Use(authMiddleware.MiddlewareFunc())
		{
			users.GET("/", ctrl.GetUsers)
			users.POST("/", ctrl.CreateUser)
			users.GET("/:id", ctrl.GetUser)
			users.PUT("/:id", ctrl.UpdateUser)
			users.DELETE("/:id", ctrl.DeleteUser)
		}

		roles := v1.Group("/roles")
		{
			roles.GET("/", ctrl.GetRoles)
			roles.GET("/:id", ctrl.GetRole)
			roles.POST("/", ctrl.PostRole)
			roles.DELETE("/:id", ctrl.DeleteRole)
		}
	}

	http.ListenAndServe(":"+port, r)
}

func ConnectMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
