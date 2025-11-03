package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddCompanyToFounder - handler to add table
func (ch *Handler) AddCompanyToFounder(c *gin.Context) {
	type addCompanyToFounder struct {
		Name string `json:"name" binfing:"required"`
	}
	request := addCompanyToFounder{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenUserID := c.GetInt64("uid")
	if err := ch.service.AddCompanyToFounder(c, request.Name, int(tokenUserID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, "Company is added to user")

}

// ReadCompany - handler to Get company info
func (ch *Handler) ReadCompany(c *gin.Context) {
	type ReadCompany struct {
		CompanyID int `json:"company_id" binding:"required"`
	}
	request := ReadCompany{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	company, err := ch.service.ReadCompany(c, request.CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, company)

}

// UpdateCompanyname - handler to update company name
func (ch *Handler) UpdateCompanyName(c *gin.Context) {
	type UpdateCompanyName struct {
		CompanyName string `json:"name" binding:"required"`
		CompanyID   int    `json:"company_id" binding:"required"`
	}

	request := UpdateCompanyName{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	tokenUserID := c.GetInt64("uid")
	if err := ch.service.UpdateCompanyName(c, request.CompanyName, request.CompanyID, int(tokenUserID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "company info updated")

}

// DeleteCompany - handler to delete company
func (ch *Handler) DeleteCompany(c *gin.Context) {
	type deleteCompany struct {
		CompanyID int `json:"company_id" binding:"required"`
	}
	request := deleteCompany{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenUserID := c.GetInt64("uid")
	res, err := ch.service.DeleteCompany(c, request.CompanyID, int(tokenUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user found for delete"})
		return
	}
	c.JSON(http.StatusOK, "company was deleted")

}

// GetCompanyID - handler to get company ID by name
func (ch *Handler) GetCompanyID(c *gin.Context) {
	type GetCompanyID struct {
		CompanyName string `json:"company_name"`
	}
	request := GetCompanyID{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := ch.service.GetCompanyID(c, request.CompanyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
