package db

import (
	"fmt"
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/config"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to the database failed:", err)
	}

	err = DB.AutoMigrate(
		&model.Problem{},
		&model.Submission{},
	)

	if err != nil {
		fmt.Printf("error while migrating %v", err.Error())
		return nil
	}
	return DB
}
