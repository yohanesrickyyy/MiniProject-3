package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbEnv := InitDB()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.Username, dbEnv.Password, dbEnv.Hostname, dbEnv.Port, dbEnv.DBName)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(fmt.Sprintf("Failed connect to database: %v", err))
	}
	fmt.Println("Success Connect to database")
	return db
}
