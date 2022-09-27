package dao

import (
	"gormexample/database"
	"gormexample/models"

	"gorm.io/gorm"
)

type EmployeeRepositoryInterface interface {
	GetAllEmployee(emp *[]models.Employee) (err error)
	CreateEmployee(emp *models.Employee) (err error)
	GetEmployeeByID(emp *models.Employee, employee_id string) (err error)
	UpdateEmployee(emp *models.Employee, employee_id string) (err error)
	DeleteEmployeeByID(emp *models.Employee, employee_id string) (err error)
}

type employeeRepo struct {
	Db *gorm.DB
}

func EmployeeNew() EmployeeRepositoryInterface {
	db := database.InitDb()
	return &employeeRepo{Db: db}
}

func (empRepo *employeeRepo) GetAllEmployee(emp *[]models.Employee) (err error) {
	if err = empRepo.Db.Find(emp).Error; err != nil {
		return err
	}

	return nil

}

func (empRepo *employeeRepo) CreateEmployee(emp *models.Employee) (err error) {

	if err = empRepo.Db.Create(emp).Error; err != nil {
		return err

	}
	return nil
}

func (empRepo *employeeRepo) GetEmployeeByID(emp *models.Employee, employee_id string) (err error) {
	if err = empRepo.Db.Where("employee_id=?", employee_id).First(emp).Error; err != nil {
		return err
	}
	return nil
}

func (empRepo *employeeRepo) UpdateEmployee(emp *models.Employee, employee_id string) (err error) {

	empRepo.Db.Where("employee_id=?", employee_id).Save(emp)
	return nil
}

func (empRepo *employeeRepo) DeleteEmployeeByID(emp *models.Employee, employee_id string) (err error) {

	empRepo.Db.Delete(emp, employee_id)

	return nil
}
