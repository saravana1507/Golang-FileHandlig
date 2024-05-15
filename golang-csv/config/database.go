package config

import (
	"fmt"
	"golang-csv/help"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Saravana1507@"
	dbName   = "user_details"
)

func DatabaseConnection() *gorm.DB {
	sqlnfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	DB, err := gorm.Open(postgres.Open(sqlnfo), &gorm.Config{})
	fmt.Println(help.Dberror(err))
	return DB
}
