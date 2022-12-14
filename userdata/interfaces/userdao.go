package interfaces

import "userdata/models"

type UserDao interface {
	Create(u *models.User) error
	/*Update(u *models.User) error
	Delete(i int) error
	GetById(i int) (models.User, error)*/
	GetAll() ([]models.User, error)
}
