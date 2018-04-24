package controllers

import (
	"fmt"
	"net/http"
	"time"

	m "savannah-go/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": projects})
}

func GetProject(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}

	var project m.Project
	projectID := c.Param("id")
	db.First(&project, projectID)

	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No roles found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": project})
}

func PostProject(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}
	var project m.Project
	db.Find(&project)

	if c.ShouldBindWith(&project, binding.JSON) == nil {
		now := time.Now()
		project.CreatedAt = now
		if len(project.Name) < 1 {

			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Name should not null",
			})

			return
		}
		db.Save(&project)

		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"message":    "projects created successfully!",
			"resourceId": project.ID,
		})
	}
}

func UpdateProject(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}
	var project m.Project
	var transformProject m.Project
	projectID := c.Param("id")

	db.First(&project, projectID)
	if c.ShouldBindWith(&transformProject, binding.JSON) == nil {
		db.Model(&project).Update(&transformProject)

		c.JSON(http.StatusCreated, gin.H{
			"status":     http.StatusCreated,
			"message":    "Project update successfully!",
			"resourceId": project,
		})
	}
}

func DeleteProject(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*gorm.DB)
	if !ok {
		fmt.Println(ok)
	}
	var project m.Project
	projectID := c.Param("id")

	db.First(&project, projectID)

	if project.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No project found!"})
		return
	}

	db.Delete(&project)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Project deleted successfully!"})
}
