package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// create function to Connect with Mysql DB
func DbConnection() *sql.DB {
	dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%v:@tcp(%v:3306)/%v?parseTime=true", user, host, dbName)

	db, err := sql.Open(dbDriver, connectionString)
	if err != nil {
		errMessage := fmt.Sprintf("error connection database : %v", err.Error())
		logrus.Error(errMessage)
		log.Fatal(errMessage)
	}

	logrus.Info("success connect to MySql")
	return db
}
