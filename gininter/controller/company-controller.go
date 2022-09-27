package controller

import (
	"gininter/entity"
	"gininter/service"

	"github.com/gin-gonic/gin"
)

type CompanyController interface {
	FindAll() []entity.Company
	Save(c *gin.Context) entity.Company
}

type controller struct {
	service service.CompanyService
}

func New(service service.CompanyService) CompanyController {
	return &controller{
		service: service,
	}
}

func (con *controller) FindAll() []entity.Company {
	return con.service.FindAll()
}

func (con *controller) Save(c *gin.Context) entity.Company {
	var company entity.Company
	c.BindJSON(&company)
	con.service.Save(company)
	return company
}
