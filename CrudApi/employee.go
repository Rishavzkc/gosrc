package main

import "gorm.io/gorm"

type Employee struct {
	gorm.Model         //This is inbulid in gorm which will auto tell time Table created,deleted
	EmpName    string  `json:"empname"`
	EmpSalary  float64 `json:"salary"`
	Email      string  `json:"email"`
}
