package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Name       string  `json:"name"`
	Id         string  `json:"id`
	Department string  `json:"dept"`
	Salary     float64 `json:"salary"`
}

var Employees = []Employee{
	{Name: "John", Id: "123", Department: "Software", Salary: 89900.98},
	{Name: "Jack", Id: "321", Department: "Accounts", Salary: 67849},
	{Name: "Andrew", Id: " 789", Department: "Admin", Salary: 56478},
}

func getEmployee(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Employees)
}

func addEmployee(c *gin.Context) {
	var newEmployee Employee
	//call BindJson to bind the received json to newEmployee
	if err := c.BindJSON(&newEmployee); err != nil {
		return
	}

	Employees = append(Employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}
func getEmployeebyId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Employees {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})

}
func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	for i, d := range Employees {
		if d.Id == id {
			Employees[i] = employee
			c.IndentedJSON(http.StatusOK, employee)
		}
	}

}

func deleteEmployee(c *gin.Context) {
	id := c.Param("id")

	for i, b := range Employees {
		if b.Id == id {

			Employees = append(Employees[:i], Employees[i+1:]...)
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Employee not found"})

}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployee)
	router.POST("/employee", addEmployee)
	router.GET("/employee/:id", getEmployeebyId)
	router.PUT("/employee/:id", updateEmployee)
	router.DELETE("/employee/:id", deleteEmployee)
	router.Run(":8090")

}
