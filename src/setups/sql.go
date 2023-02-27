package setups

import (
	"encoding/json"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupConnectionSql(option string) *gorm.DB {
	var config map[string]map[string]string
	err := json.Unmarshal([]byte(option), &config)
	if err != nil {
		panic(err)
	}
	dbUser := config["sql"]["user"]
	dbPass := config["sql"]["password"]
	dbHost := config["sql"]["host"]
	dbName := config["sql"]["db"]
	dbPort := config["sql"]["port"]
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	log.Println("Sucess to create a connection to database")
	return db
}

func CloseSqlConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
