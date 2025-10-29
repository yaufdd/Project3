package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LikeProjectPost - handler to add like on post
func (lh *Handler) LikeProjectPost(c *gin.Context) {
	type likeProjectPost struct {
		PostID int `json:"post_id" binding:"required"`
	}
	var request likeProjectPost
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := lh.service.LikeProjectPost(c, request.PostID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Post was liked")
}
