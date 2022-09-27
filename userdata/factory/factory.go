package factory

import (
	"log"
	"userdata/dao"
	"userdata/interfaces"
)

func FactoryDao(e string) interfaces.UserDao {
	var i interfaces.UserDao
	switch e {
	case "mysql":
		i = dao.UserImplMysql{}
	default:
		log.Fatalf("El motor %s no esta implementado", e)
		return nil
	}

	return i
}
