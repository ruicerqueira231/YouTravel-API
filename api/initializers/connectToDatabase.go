package initialzers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	env := os.Getenv("ENV")
	fmt.Printf("Running in %s environment\n", env)

	var dsn string
	if env == "local" {
		dsn = buildLocalDSN()
	} else {
		dsn = buildRemoteDSN()
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database in %s environment: %v", env, err)
	}
	fmt.Println("Successfully connected to the database")
}

func buildLocalDSN() string {
	host := os.Getenv("LOCAL_POSTGRES_HOST")
	user := os.Getenv("LOCAL_POSTGRES_USER")
	password := os.Getenv("LOCAL_POSTGRES_PASSWORD")
	dbname := os.Getenv("LOCAL_POSTGRES_DB")
	port := os.Getenv("LOCAL_POSTGRES_PORT")
	sslmode := os.Getenv("LOCAL_POSTGRES_SSLMODE")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC", host, user, password, dbname, port, sslmode)
}

func buildRemoteDSN() string {
	host := os.Getenv("REMOTE_POSTGRES_HOST")
	user := os.Getenv("REMOTE_POSTGRES_USER")
	password := os.Getenv("REMOTE_POSTGRES_PASSWORD")
	dbname := os.Getenv("REMOTE_POSTGRES_DB")
	port := os.Getenv("REMOTE_POSTGRES_PORT")
	sslmode := os.Getenv("REMOTE_POSTGRES_SSLMODE")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC", host, user, password, dbname, port, sslmode)
}
