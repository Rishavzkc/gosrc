package service

import (
	"gormexample/dao"
	"gormexample/models"
)

func GetAllCompany(company *[]models.Company) (err error) {

	dao.CompanyNew().GetAllCompany(company)

	if err != nil {
		return err
	}
	return nil
}

func GetCompanyByID(company *models.Company, companyid string) (err error) {

	err = dao.CompanyNew().GetCompanyByID(company, companyid)

	if err != nil {
		return err
	}

	return nil
}

func CreateCompany(company *[]models.Company) (err error) {

	err = dao.CompanyNew().CreateCompany(company)

	if err != nil {

		return err
	}

	return nil
}

func UpdateCompany(company *models.Company, companyid string) (err error) {

	err = dao.CompanyNew().UpdateCompany(company, companyid)

	if err != nil {

		return err
	}

	return nil
}

func DeleteCompany(company *models.Company, companyid string) (err error) {

	dao.CompanyNew().DeleteCompany(company, companyid)

	if err != nil {

		return err
	}

	return nil

}
