package controllers

import (
	"fmt"
	"gormexample/models"
	"gormexample/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEmployee(c *gin.Context) {

	var emp []models.Employee

	err := service.GetAllEmployee(&emp)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Employee: ": "Not found"})
	} else {
		c.JSON(http.StatusOK, emp)
	}

}

func CreateEmployee(c *gin.Context) {
	var emp models.Employee
	c.BindJSON(&emp)
	err := service.CreateEmployee(&emp)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Employee": "Is not created"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "New Employee created"})
	}
}

func GetEmployeeByID(c *gin.Context) {
	employee_id := c.Params.ByName("employee_id")
	var emp models.Employee

	err := service.GetEmployeeByID(&emp, employee_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"employee_id: " + employee_id: "Not found"})
	} else {
		c.JSON(http.StatusOK, emp)
	}
}

func UpadateEmployee(c *gin.Context) {

	var emp models.Employee
	employee_id := c.Params.ByName("employee_id")

	err := service.GetEmployeeByID(&emp, employee_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"employee_id: " + employee_id: "Not found"})
	} else {
		c.BindJSON(&emp)
		service.UpdateEmployee(&emp, employee_id)
		c.JSON(http.StatusOK, emp)
	}

}

func DeleteEmployeeByID(c *gin.Context) {
	var emp models.Employee
	employee_id := c.Params.ByName("employee_id")

	err := service.GetEmployeeByID(&emp, employee_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"employee_id: " + employee_id: "Not found"})

	} else {
		service.DeleteEmployeeByID(&emp, employee_id)
		c.JSON(http.StatusOK, gin.H{"employee_id " + employee_id: "is deleted"})
	}

}
