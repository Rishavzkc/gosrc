package main

import (
	"gininter/controller"
	"gininter/service"

	"github.com/gin-gonic/gin"
)

var (
	companyService    service.CompanyService       = service.New()
	companyController controller.CompanyController = controller.New(companyService)
)

func main() {
	server := gin.Default()

	server.GET("/company", func(c *gin.Context) {
		c.JSON(200, companyController.FindAll())
	})

	server.POST("/company", func(c *gin.Context) {
		c.JSON(200, companyController.Save(c))
	})

	server.Run(":8080")
}
