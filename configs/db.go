package configs

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DBCon : global db connection
	DBCon *gorm.DB
)

// DBInit : Initialize mysql connection
func DBInit() {

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
