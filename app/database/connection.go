package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	dbInstance *gorm.DB
	dbError    error
)

func GetDB() *gorm.DB {
	if dbInstance == nil {
		return Connect()
	}

	return dbInstance
}

func Connect() *gorm.DB {
	dbInstance, dbError = gorm.Open(postgres.Open(GetConnectionString()), &gorm.Config{})

	if dbError != nil {
		log.Fatal("DB conn: ", dbError)
	}

	return dbInstance
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSW"),
		os.Getenv("PG_DBNAME"),
	)
}
