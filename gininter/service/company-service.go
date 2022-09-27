package service

import "gininter/entity"

type CompanyService interface {
	Save(entity.Company) entity.Company
	FindAll() []entity.Company
}

type companyService struct {
	companies []entity.Company
}

func New() CompanyService {
	return &companyService{}
}

func (service *companyService) Save(company entity.Company) entity.Company {
	service.companies = append(service.companies, company)
	return company
}

func (service *companyService) FindAll() []entity.Company {
	return service.companies
}
