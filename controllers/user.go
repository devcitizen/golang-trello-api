package controllers

import (
	"fmt"
	"net/http"
	m "savannah-go/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/validator.v2"
)

var conn = "root:12345678@/savannah-go?charset=utf8&parseTime=True&loc=Local"

type (
	users struct {
		ID        uint   `json:"id"`
		Name      string `validate:"nonzero"`
		Email     string `gorm:"type:varchar(100);unique_index"`
		RoleID    int    `validate:"nonzero"`
		Address   string `validate:"nonzero"`
		CreatedAt time.Time
	}
)

func GetUsers(c *gin.Context) {
	db, err := gorm.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	var users []m.User
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No users found!"})
		return
	}

	for i, _ := range users {
		db.Model(users[i]).Related(&users[i].Role)
	}

	//transforms the todos for building a good response
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func CreateUser(c *gin.Context) {
	db, err := gorm.Open("mysql", conn)
	var user users
	var role m.Role
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	if c.ShouldBindWith(&user, binding.JSON) == nil {
		if errs := validator.Validate(user); errs != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusBadRequest,
				"message": errs,
			})
			return
		}

		db.First(&role, user.RoleID)
		now := time.Now()
		user.CreatedAt = now
		if role.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Cannot find role!"})
			return
		}

		db.Save(&user)

		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"message":    "User created successfully!",
			"resourceId": user.ID,
		})
	}
}

func GetUser(c *gin.Context) {
	db, err := gorm.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	var user m.User
	userID := c.Param("id")

	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func UpdateUser(c *gin.Context) {
	db, err := gorm.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	var user m.User
	var transformuser users
	userID := c.Param("id")

	db.First(&user, userID)
	if c.ShouldBindWith(&transformuser, binding.JSON) == nil {
		db.Model(&user).Update(&transformuser)

		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"message":    "User update successfully!",
			"resourceId": user,
		})
	}
}

func DeleteUser(c *gin.Context) {
	db, err := gorm.Open("mysql", conn)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	var user m.User
	userID := c.Param("id")

	db.First(&user, userID)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
