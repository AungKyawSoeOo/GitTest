package config

import (
	"fmt"
	"gin_test/bulletin_board/helper"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	helper.ErrorPanic(err)
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	DNS := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Database Connected!")

	return db
}
