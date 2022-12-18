package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWD")
	// dbAddress := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")
	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbAddress, dbPort, dbName)
	// db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Database Connection failed %v",err.Error())
	// }
	// return db

	dsn := "host=localhost user=postgres password=fitri123 dbname=gogrpc port=5432"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Database Connection failed %v",err.Error())
	}
	return db
}