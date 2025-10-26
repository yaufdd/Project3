package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
)

// AddComment - handler to add comment to project post
func (cmh *Handler) AddComment(c *gin.Context) {
	var comment *models.Comment
	if err := c.ShouldBindBodyWithJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cmh.service.AddComment(c, comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errir": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user published comment to post"})
}
