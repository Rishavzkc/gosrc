package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeID           uint   `json:"employeeid"`
	EmployeeName         string `json:"employeename"`
	EmployeeMobileNumber string `json:"employeemobilenumber"`
	CompanyID            uint   `json:"companyid"`
}

func (e *Employee) TableName() string {
	return "employee"
}
