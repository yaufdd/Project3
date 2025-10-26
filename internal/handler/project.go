package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

// AddProject = handler to add project into table
func (ph *Handler) AddProject(c *gin.Context) {
	var newProject *models.Project
	if err := c.ShouldBindJSON(&newProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ph.service.AddProject(c, newProject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "project was added")
}

// ReadProject - handler to get project info
func (ph *Handler) ReadProject(c *gin.Context) {
	type idForRead struct{ ID int }
	var id idForRead
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	project, err := ph.service.ReadProject(c, id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

// UpdateProjectTitle - handler to update project title
func (ph *Handler) UpdateProjectTitle(c *gin.Context) {
	var newProjectInfo *models.Project
	if err := c.ShouldBindJSON(&newProjectInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ph.service.UpdateProjectTitle(c, newProjectInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "project info was updated")
}

// DeleteProject - handler to delete project from table
func (ph *Handler) DeleteProject(c *gin.Context) {
	type idForDelete struct{ ID int }
	var id idForDelete
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := ph.service.DeleteProject(c, id.ID)
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
	var request request.GetProjectID
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errir": err.Error()})
		return
	}
	res, err := ph.service.GetProjectID(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
