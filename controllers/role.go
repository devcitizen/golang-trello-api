package controllers

import (
	"fmt"
	"net/http"
	m "savannah-go/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

func GetRoles(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var roles []m.Role
	db.Find(&roles)

	if len(roles) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No roles found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": roles})
}

func GetRole(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var role m.Role
	rolesID := c.Param("id")
	db.First(&role, rolesID)

	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No roles found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": role})
}

func PostRole(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}
	var role m.Role

	if c.ShouldBindWith(&role, binding.JSON) == nil {
		now := time.Now()
		role.CreatedAt = now
		if len(role.Name) < 1 {

			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Name should not null",
			})

			return
		}
		db.Save(&role)

		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"message":    "Role created successfully!",
			"resourceId": role.ID,
		})
	}
}

func DeleteRole(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}
	var role m.Role
	roleID := c.Param("id")

	db.First(&role, roleID)

	if role.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No role found!"})
		return
	}

	db.Delete(&role)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Role deleted successfully!"})
}
