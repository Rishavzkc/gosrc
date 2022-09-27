package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyID    uint       `json:"companyid"`
	CompanyName  string     `json:"companyename"`
	CompanyEmail string     `json:"companyemail"`
	Employees    []Employee `gorm:"foreignkey:company_id"`
}

func (c *Company) TableName() string {
	return "company"
}
