package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID           uint
	EmployeeName string
}

type Company struct {
	gorm.Model
	ID        uint
	Name      string
	Location  string
	Employees []Employee `gorm:"foreignkey:company_id"`
}

func createCompany() {

	employees := []Employee{
		{EmployeeName: "John"},
		{EmployeeName: "Jack"},
		{EmployeeName: "Andrew"},
		{EmployeeName: "Tom"},
		{EmployeeName: "Ram"},
	}

	for i, _ := range employees {
		if err := db.Create(&employees[i]).Error; err != nil {
			log.Fatal(err)
		}
	}

	companies := []Company{
		{Name: "Quest", Location: "Pune", Employees: []Employee{employees[0], employees[2]}},
		{Name: "TCS", Location: "Blr", Employees: []Employee{employees[0], employees[1]}},
		{Name: "Wipro", Location: "Delhi", Employees: []Employee{employees[3]}},
		{Name: "Infosys", Location: "Noida", Employees: []Employee{employees[2]}},
		{Name: "Tr", Location: "Pune", Employees: []Employee{employees[4], employees[1]}},
	}

	for i, _ := range companies {
		if err := db.Create(&companies[i]).Error; err != nil {
			log.Fatal(err)
		}
	}
}

var db *gorm.DB

func connection() {
	var err error
	db, err = gorm.Open(mysql.Open("root:Quest1234@tcp(localhost:3306)/gorm?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Company{}, &Employee{})

}

func GetAllCompanay() ([]Company, error) {
	var companies []Company
	err := db.Model(&Company{}).Preload("Employees").Find(&companies).Error
	return companies, err

}

func main() {
	connection()

	var input int

	for {
		fmt.Println("Enter the Details")
		fmt.Println("1 For Create")
		fmt.Println("2 For Read")
		fmt.Println("3 For Update")
		fmt.Println("4 For Delete")
		fmt.Println("5 For Exit")

		fmt.Scanln(&input)

		switch input {
		case 1:
			createCompany()
		case 2:
			GetAllCompanay()
		case 3:
			//updateCompany(Employees)
		case 4:
			//deleteCompany(Employees)
		default:
			fmt.Println("Exit")
		}

	}

}
