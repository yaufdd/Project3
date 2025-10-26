package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

// AddCompanyToFounder - handler to add table
func (ch *Handler) AddCompanyToFounder(c *gin.Context) {
	var newCompany *models.Company
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ch.service.AddCompanyToFounder(c, newCompany); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, "Company is added to user")

}

// ReadCompany - handler to Get company info
func (ch *Handler) ReadCompany(c *gin.Context) {
	type idForRead struct{ ID int }
	var id idForRead
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	company, err := ch.service.ReadCompany(c, id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, company)

}

// UpdateCompanyname - handler to update company name
func (ch *Handler) UpdateCompanyname(c *gin.Context) {
	var newCompanyInfo *models.Company
	if err := c.ShouldBindJSON(&newCompanyInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := ch.service.UpdateCompanyname(c, newCompanyInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "company info updated")

}

// DeleteCompany - handler to delete company
func (ch *Handler) DeleteCompany(c *gin.Context) {
	type idForDelete struct{ ID int }
	var id idForDelete
	if err := c.ShouldBindJSON(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := ch.service.DeleteCompany(c, id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if res == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no user found for delete"})
	}
	c.JSON(http.StatusOK, "company was deleted")

}

// GetCompanyID - handler to get company ID by name
func (ch *Handler) GetCompanyID(c *gin.Context) {
	var request request.GetCompanyID
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	res, err := ch.service.GetCompanyID(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
