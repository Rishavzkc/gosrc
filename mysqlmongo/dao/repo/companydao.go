package repo

import "mysqlmongo/model"

type CompanyDao interface {
	CreateCompany(*model.Company) error
	GetCompany(string) (*model.Company, error)
	GetAll() ([]*model.Company, error)
	UpdateCompany(*model.Company) error
	DeleteCompany(string) error
}
