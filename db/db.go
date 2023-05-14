package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tirthankarkundu17/ecommerce-price-checker/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func NewDB() (*DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	schema := os.Getenv("DB_SCHEMA")
	dsn := username + ":" + password + "@tcp(" + host + ")/" + schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	db := &DB{
		Conn: conn,
	}

	return db, nil
}

func (db *DB) Close() {
	sqlDB, err := db.Conn.DB()
	if err != nil {
		fmt.Println("Error closing database:", err)
		return
	}
	sqlDB.Close()
}

func (db *DB) AutoMigrate() error {
	return db.Conn.AutoMigrate(&model.UserProduct{}, &model.User{})
}
