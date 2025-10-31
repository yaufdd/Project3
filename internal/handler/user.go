// Package handler includes all function for handler
package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
)

// CRUD for user
// CreatUser - handler for creating user.
// Used for authorization
func (uh *Handler) CreateUser(c *gin.Context) {
	var request *models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, _, err := uh.service.CreateUser(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User is saved to db")

}

// ReadUser - handler for get user info by ID
func (uh *Handler) ReadUser(c *gin.Context) {
	type readUser struct {
		UserID int `json:"user_id" binding:"required"`
	}
	request := readUser{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := uh.service.ReadUser(c, request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser - handler for update user. Need to user id
func (uh *Handler) UpdateUsername(c *gin.Context) {
	type updateUsername struct {
		UserID      int    `json:"user_id" binding:"required"`
		NewUsername string `json:"new_uname" binding:"required"`
	}
	request := updateUsername{}
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uh.service.UpdateUsername(c, request.UserID, request.NewUsername); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "user updated")
}

// DeleteUser - handler for delete user by id
func (uh *Handler) DeleteUser(c *gin.Context) {
	type deleteUser struct {
		UserID int
	}
	request := deleteUser{}
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := uh.service.DeleteUser(c, request.UserID)
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
	type getUserID struct {
		Username string `json:"username" binding:"required"`
	}
	request := getUserID{}
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errir": err.Error()})
		return
	}
	res, err := uh.service.GetUserID(c, request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// RegistrationJWT registration with jwt token
func (uh *Handler) Registration(c *gin.Context) {
	var request *models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, err := uh.service.Registration(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("{access_token : %s}\n{refresh_token: %s}", accessToken, refreshToken))
}

func (uh *Handler) Authentication(c *gin.Context) {
	type AuthCredential struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	request := AuthCredential{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := uh.service.Authentication(c, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("{access_token : %s}", accessToken))
}
