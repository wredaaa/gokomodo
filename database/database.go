package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func GetDatabase() *gorm.DB {

	// Ex: user:password@tcp(host:port)/dbname
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PROTOCOL"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	connection, err := gorm.Open("mysql", mysqlCredentials)

	sqldb := connection.DB()
	err = sqldb.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection successful.")
	return connection
}

/*
	Do this function to close database connection.

*
*/
func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}
