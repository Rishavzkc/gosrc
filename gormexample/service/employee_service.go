package service

import (
	"gormexample/dao"
	"gormexample/models"
)

func GetAllEmployee(emp *[]models.Employee) (err error) {

	err = dao.EmployeeNew().GetAllEmployee(emp)
	if err != nil {
		return err
	}

	return nil

}

func CreateEmployee(emp *models.Employee) (err error) {

	err = dao.EmployeeNew().CreateEmployee(emp)
	if err != nil {
		return err

	}
	return nil
}

func GetEmployeeByID(emp *models.Employee, employee_id string) (err error) {

	err = dao.EmployeeNew().GetEmployeeByID(emp, employee_id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateEmployee(emp *models.Employee, employee_id string) (err error) {

	err = dao.EmployeeNew().UpdateEmployee(emp, employee_id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmployeeByID(emp *models.Employee, employee_id string) (err error) {

	err = dao.EmployeeNew().DeleteEmployeeByID(emp, employee_id)
	if err != nil {
		return err
	}

	return nil
}
