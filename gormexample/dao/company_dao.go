package dao

import (
	"fmt"
	"gormexample/database"
	"gormexample/models"

	"gorm.io/gorm"
)

type CompanyService interface {
	GetAllCompany(company *[]models.Company) (err error)
	GetCompanyByID(company *models.Company, companyid string) (err error)
	CreateCompany(company *[]models.Company) (err error)
	UpdateCompany(company *models.Company, companyid string) (err error)
	DeleteCompany(company *models.Company, company_id string) (err error)
}

type CompanyRepo struct {
	Db *gorm.DB
}

func CompanyNew() CompanyService {
	db := database.InitDb()

	return &CompanyRepo{Db: db}
}

func (companyRepo *CompanyRepo) GetAllCompany(company *[]models.Company) (err error) {

	if err = companyRepo.Db.Find(company).Error; err != nil {
		return err

	}

	return nil
}

func (companyRepo *CompanyRepo) GetCompanyByID(company *models.Company, companyid string) (err error) {

	if err = companyRepo.Db.Preload("Employee").Where("company_id = ?", companyid).First(company).Error; err != nil {
		return err

	}

	return nil
}

func (companyRepo *CompanyRepo) CreateCompany(company *[]models.Company) (err error) {

	for _, companies := range *company {

		companyRepo.Db.Create(&companies)
	}

	if err != nil {

		return err
	}

	return nil
}

func (companyRepo *CompanyRepo) UpdateCompany(company *models.Company, companyid string) (err error) {

	fmt.Println(company)
	companyRepo.Db.Save(company)

	return nil
}

func (companyRepo *CompanyRepo) DeleteCompany(company *models.Company, company_id string) (err error) {

	companyRepo.Db.Where("company_id = ?", company_id).Delete(company)

	return nil

}
