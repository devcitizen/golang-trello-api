package controllers

import (
	"fmt"
	"net/http"
	m "savannah-go/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/validator.v2"
)

func GetSprints(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var sprints []m.Sprint
	projectID := c.Param("id")
	db.Where("project_id = ?", projectID).Order("id desc").Find(&sprints)

	if len(sprints) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Sprints found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": sprints})
}

func GetSprintBacklog(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var sprint m.Sprint
	sprintID := c.Param("id")
	db.First(&sprint, sprintID).Related(&sprint.Backlog)

	if len(sprint.Backlog) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No Backlog found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": sprint})
}

func CreateSprint(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var sprint m.Sprint
	var project m.Project
	if c.ShouldBindWith(&sprint, binding.JSON) == nil {
		if errs := validator.Validate(sprint); errs != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusBadRequest,
				"message": errs,
			})
			return
		}

		if err := db.First(&project, sprint.ProjectID).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusNotFound, "message": "Cannot find project!"})
			return
		}

		db.Save(&sprint)

		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"message":    "Sprint created successfully!",
			"resourceId": sprint.ID,
		})
	}
}
