package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// loading env vars
	if err := godotenv.Load(); err != nil {
		panic("Env file not found")
	}

	// loading db env vars
	pgDBName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil, fmt.Errorf("can not load environmental var POSTGRES_DB")
	}

	pgUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, fmt.Errorf("can not load environmental var POSTGRES_USER")
	}

	pgPwd, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("can not load environmental var POSTGRES_PASSWORD")
	}

	pgHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return nil, fmt.Errorf("can not load environmental var POSTGRES_HOST")
	}

	pgPort, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return nil, fmt.Errorf("can not load environmental var POSTGRES_PORT")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		pgHost,
		pgUser,
		pgPwd,
		pgDBName,
		pgPort,
	)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can not connect to DB")
	}

	return db, nil
}
