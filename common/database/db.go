package database

import (
	"final-project/user/entity/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDb   = "finalproject"
	PostgresUser = "postgres"
	PostgresPass = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgresUser, PostgresPass, PostgresDb,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err := db.AutoMigrate(model.User{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
