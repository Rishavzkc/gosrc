package factory

import (
	"log"
	"mysqlmongo/dao/mysql"
	"mysqlmongo/dao/repo"
)

func FactoryDao(e string) repo.CompanyDao {
	var i repo.CompanyDao
	switch e {
	case "mysql":
		i = &mysql.CompanyRepoImpl{}
	default:
		log.Fatalf("No connection found ", e)
		return nil
	}

	return i
}
