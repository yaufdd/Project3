package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/request"
)

// LikeProjectPost - handler to add like on post
func (lh *Handler) LikeProjectPost(c *gin.Context) {
	var likeRequest request.LikePost
	if err := c.ShouldBindJSON(&likeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := lh.service.LikeProjectPost(c, likeRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Post was liked")
}
