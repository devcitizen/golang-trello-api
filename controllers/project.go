package controllers

import (
	"fmt"
	"net/http"
	m "savannah-go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetProjects(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var projects []m.Project
	db.Find(&projects)

	if len(projects) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Projects found!"})
		return
	}
}

func GetProject(c *gin.Context) {

}

func PostProject(c *gin.Context) {

}

func UpdateProject(c *gin.Context) {

}

func DeleteProject(c *gin.Context) {

}
