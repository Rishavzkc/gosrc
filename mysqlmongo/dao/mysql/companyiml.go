package mysql

import (
	"fmt"

	"mysqlmongo/dao/repo"
	"mysqlmongo/model"

	"gorm.io/gorm"
)

type CompanyRepoImpl struct {
	companycollection *gorm.DB
}

func NewCompanyService(companycollection *gorm.DB) repo.CompanyDao {
	return &CompanyRepoImpl{
		companycollection: companycollection,
	}
}

func (c *CompanyRepoImpl) CreateCompany(company *model.Company) error {
	if tx := c.companycollection.Create(company); tx.Error != nil {
		return fmt.Errorf("failed to insert record in company db: %w", tx.Error)
	}
	return nil

}

func (c *CompanyRepoImpl) GetCompany(id string) (*model.Company, error) {

	var company model.Company
	if tx := c.companycollection.First(&company, "id = ?", id); tx.Error != nil {
		return nil, fmt.Errorf("failed to fetch record with id '%s': %w", id, tx.Error)
	}
	return &company, nil

}

func (c *CompanyRepoImpl) GetAll() ([]*model.Company, error) {

	var companies []*model.Company
	if tx := c.companycollection.Find(&companies); tx.Error != nil {
		return nil, fmt.Errorf("failed to fetch all records from company db: %w", tx.Error)
	}
	return companies, nil

}

func (c *CompanyRepoImpl) UpdateCompany(company *model.Company) error {

	tx := c.companycollection.Model(model.Company{ID: company.ID}).Updates(&model.Company{
		Name:     company.Name,
		Location: company.Location,
	})
	if tx.Error != nil {
		return fmt.Errorf("failed to update record with id '%s' in company db: %w", company.ID, tx.Error)
	}
	return nil
}

func (c *CompanyRepoImpl) DeleteCompany(id string) error {
	tx := c.companycollection.Delete(&model.Company{ID: id})
	if tx.Error != nil {
		return fmt.Errorf("failed to delete record with id '%s': %w", id, tx.Error)
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("no record found to delete for id '%s'", id)
	}
	return nil

}
