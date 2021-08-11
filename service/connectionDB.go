package service

import (
	"apigolang/helpers"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConnection(database string) {

}

func NewConnection() *gorm.DB {
	dbUser := helpers.GetEnvParamDefault("DB_USER", "root")
	dbPass := helpers.GetEnvParamDefault("DB_PASS", "")
	dbHost := helpers.GetEnvParamDefault("DB_HOST", "127.0.0.1")
	dbPort := helpers.GetEnvParamDefault("DB_PORT", "3306")
	dbDatabase := helpers.GetEnvParam("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDatabase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.HandleError(err, "Failed to make connection to database")
	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	helpers.HandleError(err, "Failed to close database connection:")
	dbSQL.Close()
}
