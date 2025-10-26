package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

// PublishProjectPost - handler for publish project post
func (psth *Handler) PublishProjectPost(c *gin.Context) {
	var publishReq *models.ProjectPost
	if err := c.ShouldBindJSON(&publishReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := psth.service.PublishProjectPost(c, publishReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, "post was published")
}

// ReadProjectPost a handler to get project post info
func (psth *Handler) ReadProjectPost(c *gin.Context) {
	type idForProjectPost struct{ ID int }
	var id idForProjectPost
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	post, err := psth.service.ReadProjectPost(c, id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdateProjectPostDescription - handler to update project post description
func (psth *Handler) UpdateProjectPostDescription(c *gin.Context) {
	var newDesc request.NewDescription
	if err := c.ShouldBindJSON(&newDesc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := psth.service.UpdateProjectPostDescription(c, newDesc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, "Post was updated")
}

// DeleteProjectPost - handler to delete project post
func (psth *Handler) DeleteProjectPost(c *gin.Context) {
	type idForDelete struct{ ID int }
	var id idForDelete
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	res, err := psth.service.DeleteProjectPost(c, id.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if res == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User was found but was not deleted"})
	}
	c.JSON(http.StatusOK, "Post was deleted")
}
