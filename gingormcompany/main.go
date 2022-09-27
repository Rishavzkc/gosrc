package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	var urlDSN = "root:Quest1234@tcp(localhost:3306)/newdb?parseTime=true"
	db, err = gorm.Open(mysql.Open(urlDSN))
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&Company{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/company")
	{
		v1.POST("/", createCompany)
		v1.GET("/", getAllCompany)
		v1.GET("/:id", getCompanyById)
		v1.PUT("/:id", updateCompany)
		v1.DELETE("/:id", deleteCompany)
	}
	router.Run()

}

type (
	Company struct {
		gorm.Model

		Name   string `json:"name"`
		Status int    `json:"status"`
	}
	newCompany struct {
		Id     uint   `json:"id"`
		Name   string `json:"name"`
		Status bool   `json:"status"`
	}
)

func createCompany(c *gin.Context) {
	status, _ := strconv.Atoi(c.PostForm("status"))
	company := Company{Name: c.PostForm("name"), Status: status}
	db.Save(&company)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Company details created successfully!", "resourceId": company.ID})
}

func getAllCompany(c *gin.Context) {
	var company []Company
	var _company []newCompany

	db.Find(&company)

	if len(company) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
		return
	}

	for _, item := range company {
		status := false
		if item.Status == 1 {
			status = true
		} else {
			status = false
		}
		_company = append(_company, newCompany{Id: item.ID, Name: item.Name, Status: status})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _company})
}

func getCompanyById(c *gin.Context) {
	var company Company
	companyID := c.Param("id")

	db.First(&company, companyID)
	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
		return
	}
	status := false
	if company.Status == 1 {
		status = true
	} else {
		status = false
	}
	_company := newCompany{Id: company.ID, Name: company.Name, Status: status}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _company})
}

func updateCompany(c *gin.Context) {

	var company Company
	companyID := c.Param("id")

	db.First(&company, companyID)

	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
		return
	}

	db.Model(&company).Update("Name", c.PostForm("Name"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	db.Model(&company).Update("completed", status)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Company updated successfully!"})
}
func deleteCompany(c *gin.Context) {
	var company Company
	companyID := c.Param("id")

	db.First(&company, companyID)

	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
		return
	}

	db.Delete(&company)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Company deleted successfully!"})
}
