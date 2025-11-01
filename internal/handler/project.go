package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
)

// AddProject = handler to add project into table
func (ph *Handler) AddProject(c *gin.Context) {
	var request *models.Project
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ph.service.AddProject(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "project was added")
}

// ReadProject - handler to get project info
func (ph *Handler) ReadProject(c *gin.Context) {
	type readProject struct {
		ProjectID int `json:"project_id" binding:"required"`
	}
	request := readProject{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := ph.service.ReadProject(c, request.ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

// UpdateProjectTitle - handler to update project title
func (ph *Handler) UpdateProjectTitle(c *gin.Context) {
	type updateProjectTitle struct {
		ProjectID int    `json:"project_id" binding:"required"`
		Title     string `json:"title" binding:"required"`
		UserID    int    `json:"user_id" binding:"required"`
	}
	request := updateProjectTitle{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ph.service.UpdateProjectTitle(c, request.ProjectID, request.Title); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "project info was updated")
}

// DeleteProject - handler to delete project from table
func (ph *Handler) DeleteProject(c *gin.Context) {
	type deleteProject struct {
		ProjectID int `json:"project_id" binding:"required"`
		UserID    int `json:"user_id" binding:"required"`
	}
	request := deleteProject{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := ph.service.DeleteProject(c, request.ProjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if res == 0 {
		c.JSON(http.StatusInternalServerError, "project was not deleted")
	}
	c.JSON(http.StatusOK, "project was deleted")
}

// GetProjectID - handler to get ID of project
func (ph *Handler) GetProjectID(c *gin.Context) {
	type getProjectID struct {
		ProjectTitle string `json:"project_title" binding:"required"`
	}
	request := getProjectID{}
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errir": err.Error()})
		return
	}
	res, err := ph.service.GetProjectID(c, request.ProjectTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
