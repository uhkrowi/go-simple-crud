package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetEnv(key string) string {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	DBConn *gorm.DB
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func getDBConfig() (config DBConfig) {
	config.Host = GetEnv("DB_HOST")
	config.User = GetEnv("DB_USER")
	config.Password = GetEnv("DB_PASSWORD")
	config.Port = GetEnv("DB_PORT")
	config.DBName = GetEnv("DB_NAME")

	return
}

func InitDB() {
	config := getDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DBConn = db
}

// func CloseDBConnection(connection *gorm.DB) {
func CloseDBConnection() {
	conn, _ := DBConn.DB()
	conn.Close()
	fmt.Println("sip, db is closed")
}
