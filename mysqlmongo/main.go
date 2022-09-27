package main

import (
	"log"
	"mysqlmongo/controllers"
	"mysqlmongo/dao/mysql"
	"mysqlmongo/dao/repo"
	"mysqlmongo/factory"
	"mysqlmongo/utilities"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                      = mysql.SetupDatabaseConnection()
	companyService    repo.CompanyDao               = mysql.NewCompanyService(db)
	companyController controllers.CompanyController = controllers.New(companyService)
)

func main() {
	//	defer mysql.CloseDatabaseConnection(db)
	r := gin.Default()

	companyRoutes := r.Group("/company")
	{
		companyRoutes.GET("/", companyController.GetAll)
		companyRoutes.POST("/", companyController.CreateCompany)
		companyRoutes.GET("/:id", companyController.GetCompany)
		companyRoutes.PUT("/:id", companyController.UpdateCompany)
		companyRoutes.DELETE("/:id", companyController.DeleteCompany)
	}

	r.Run()

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	factory.FactoryDao(config.Engine)

	if err != nil {
		log.Fatal(err)
		return
	}

}
