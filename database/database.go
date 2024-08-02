package database

import (
	"fmt"
	"log"
	"os"
	"todoAPI/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct{
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		
        Logger: logger.Default.LogMode(logger.Info),
	})

    if err != nil {
        log.Fatal("Failed to connect to the database! \n", err)
        os.Exit(2)
    }

    fmt.Println("Database connected!")
    db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")

    db.AutoMigrate(&models.Todo{})

	DB = Dbinstance{
		Db: db,
	}
}
