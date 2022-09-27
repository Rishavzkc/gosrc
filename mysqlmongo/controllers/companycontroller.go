package controllers

import (
	"mysqlmongo/dao/repo"
	"mysqlmongo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	CompanyService repo.CompanyDao
}

func New(companyservice repo.CompanyDao) CompanyController {
	return CompanyController{
		CompanyService: companyservice,
	}
}

func (cc *CompanyController) CreateCompany(ctx *gin.Context) {
	var company model.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CompanyService.CreateCompany(&company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (cc *CompanyController) GetCompany(ctx *gin.Context) {
	companyname := ctx.Param("id")
	company, err := cc.CompanyService.GetCompany(companyname)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, company)

}

func (cc *CompanyController) GetAll(ctx *gin.Context) {
	companies, err := cc.CompanyService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

func (cc *CompanyController) UpdateCompany(ctx *gin.Context) {
	var company model.Company
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := cc.CompanyService.UpdateCompany(&company)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
	//ctx.JSON(http.StatusOK, gin.H{"data": company})
}

func (cc *CompanyController) DeleteCompany(ctx *gin.Context) {
	companyid := ctx.Param("id")
	err := cc.CompanyService.DeleteCompany(companyid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success"})
}
