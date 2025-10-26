// Package handler includes all function for handler
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

// CreatUser - handler for creating user.
// Used for authorization
func (uh *Handler) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.service.CreateUser(c, &newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User is saved to db")

}

// ReadUser - handler for get user info by ID
func (uh *Handler) ReadUser(c *gin.Context) {
	type ifForRead struct{ ID int }
	var r ifForRead
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uh.service.ReadUser(c, r.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser - handler for update user. Need to send full user json object
func (uh *Handler) UpdateUsername(c *gin.Context) {
	var newUserInfo *models.User
	if err := c.ShouldBindBodyWithJSON(&newUserInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.service.UpdateUsername(c, newUserInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "user updated")
}

// DeleteUser - handler for delete user by id
func (uh *Handler) DeleteUser(c *gin.Context) {
	type idForDelete struct{ ID int }
	var id idForDelete
	if err := c.ShouldBindBodyWithJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := uh.service.DeleteUser(c, id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user found for delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// GetUserID - handler to get user ID
func (uh *Handler) GetUserID(c *gin.Context) {
	var request request.GetUserID
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errir": err.Error()})
		return
	}
	res, err := uh.service.GetUserID(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
