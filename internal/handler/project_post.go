package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
)

// PublishProjectPost - handler for publish project post
func (psth *Handler) PublishProjectPost(c *gin.Context) {
	var request *models.ProjectPost
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := psth.service.PublishProjectPost(c, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, "post was published")
}

// ReadProjectPost a handler to get project post info
func (psth *Handler) ReadProjectPost(c *gin.Context) {
	type readProjectPost struct {
		PostID int `json:"post_id" binding:"required"`
	}
	request := readProjectPost{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	post, err := psth.service.ReadProjectPost(c, request.PostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, post)
}

// UpdateProjectPostDescription - handler to update project post description
func (psth *Handler) UpdateProjectPostDescription(c *gin.Context) {
	type updateProjectPostDescription struct {
		PostID      int    `json:"post_id" binding:"required"`
		Description string `json:"desc" binding:"required"`
		UserID      int    `json:"user_id" binding:"required"`
	}
	request := updateProjectPostDescription{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := psth.service.UpdateProjectPostDescription(c, request.PostID, request.Description)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, "Post was updated")
}

// DeleteProjectPost - handler to delete project post
func (psth *Handler) DeleteProjectPost(c *gin.Context) {
	type deleteProjectPost struct {
		PostID int
		UserID int `json:"user_id" binding:"required"`
	}
	request := deleteProjectPost{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	res, err := psth.service.DeleteProjectPost(c, request.PostID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if res == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User was found but was not deleted"})
	}
	c.JSON(http.StatusOK, "Post was deleted")
}
