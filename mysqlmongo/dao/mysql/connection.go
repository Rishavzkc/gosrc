package mysql

import (
	"fmt"
	"mysqlmongo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*func get() *sql.DB {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
} */

func SetupDatabaseConnection() *gorm.DB {

	var urlDSN = "root:Quest1234@tcp(localhost:3306)/newdb?parseTime=true"
	var err error

	db, err := gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&model.Company{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
