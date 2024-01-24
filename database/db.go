package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/prashant42b/crud-task/config"
	"github.com/prashant42b/crud-task/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	// p := config.Config("DB_PORT")
	p := config.PORT
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.HOST, port, config.USER, config.PASSWORD, config.NAME)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.CRUD{})
	fmt.Println("Database Migrated")
}
