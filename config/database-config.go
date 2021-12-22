package config

import (
	"fmt"
	"golang_api_todo_books/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabaseConnection
func SetupDatabaseConnection() *gorm.DB {
	// errEnv := godotenv.Load()

	// if errEnv != nil {
	// 	panic("Failed to load env file ")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s sslmode=require TimeZone=Asia/Shanghai", dbUser, dbPass, dbPort, dbHost, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database ")
	}

	// nanti kita isi modelnya disini
	db.AutoMigrate(&entity.Book{}, &entity.User{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection to database")
	}

	dbSQL.Close()
}
