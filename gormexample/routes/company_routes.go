package routes

import (
	"gormexample/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	grp := r.Group("/company-api/")
	{

		grp.GET("company", controllers.GetAllComp)
		grp.POST("company", controllers.CreateComp)
		grp.GET("company/:companyid", controllers.GetCompanyByIDCon)
		grp.PUT("company/:companyid", controllers.UpdateCompanyCon)
		grp.DELETE("company/:companyid", controllers.DeleteCompanyCon)

		grp.GET("employee", controllers.GetAllEmployee)
		grp.POST("employee", controllers.CreateEmployee)
		grp.GET("employee/:employee_id", controllers.GetEmployeeByID)
		grp.PUT("employee/:employee_id", controllers.UpadateEmployee)
		grp.DELETE("employee/:employee_id", controllers.DeleteEmployeeByID)

	}

	return r
}
