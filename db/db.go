package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DBCon : global db connection
	DBCon *gorm.DB
)

// Init : Initialize mysql connection
func Init() {
	envver := godotenv.Load()
	if envver != nil {
		log.Fatal("Error loading .env file")
	}

	DBUser := os.Getenv("DB_USER")
	DBPasswd := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	ConnectionString := "tcp(" + DBHost + ":" + DBPort + ")"

	fmt.Printf("%s:%s@%s/%s?charset=utf8&parseTime=True\n", DBUser, DBPasswd, ConnectionString, DBName)

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPasswd, ConnectionString, DBName)

	var err error
	DBCon, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}
