package controllers

import (
	"fmt"
	"gormexample/models"

	"net/http"
	"gormexample/service"
	"github.com/gin-gonic/gin"
)

func GetAllComp(c *gin.Context) {

	fmt.Println("In controller start")
	var company []models.Company

	err := service.GetAllCompany(&company)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Company: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, company)
	}

	fmt.Println("In controller end")
}

func CreateComp(c *gin.Context) {
	var company []models.Company
	errs := c.BindJSON(&company)

	if errs != nil {
		return
	}

	err := service.CreateCompany(&company)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"Company": "Is not created"})
		return
	} else {
		c.JSON(http.StatusCreated, company)
		return
	}
}

func GetCompanyByIDCon(c *gin.Context) {
	companyid := c.Params.ByName("companyid")
	var company models.Company

	err := service.GetCompanyByID(&company, companyid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CompanyID: " + companyid: "Not found"})
	} else {
		c.JSON(http.StatusOK, company)
	}

}

func UpdateCompanyCon(c *gin.Context) {

	var company models.Company
	companyid := c.Params.ByName("companyid")

	err := service.GetCompanyByID(&company, companyid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CompanyID: " + companyid: "Not found"})
	} else {
		c.BindJSON(&company)
		err = service.UpdateCompany(&company, companyid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CompanyID: " + companyid: "Not found"})
		}
		c.JSON(http.StatusOK, company)
	}
}

func DeleteCompanyCon(c *gin.Context) {
	var company models.Company
	companyid := c.Params.ByName("companyid")

	err := service.GetCompanyByID(&company, companyid)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"CompanyID: " + companyid: "Not found"})
	} else {

		service.DeleteCompany(&company, companyid)

		c.JSON(http.StatusOK, gin.H{"company_id " + companyid: "is deleted"})

	}
}
