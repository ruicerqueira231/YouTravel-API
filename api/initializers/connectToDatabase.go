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
	fmt.Printf("Running in %s environment\n", os.Getenv("ENV"))

	if env == "local" {
		connectToLocalDB()
	} else {
		connectToRemoteDB()
	}
}

func connectToRemoteDB() {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "ep-small-fire-a202x97i-pooler.eu-central-1.aws.neon.tech"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		user = "default"
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password == "" {
		password = "kZJGLz6Xwl4d"
	}
	dbname := os.Getenv("POSTGRES_DATABASE")
	if dbname == "" {
		dbname = "verceldb"
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		port = "5432"
	}
	sslmode := os.Getenv("POSTGRES_SSLMODE")
	if sslmode == "" {
		sslmode = "require"
	}

	fmt.Printf("Connecting to remote DB with host=%s user=%s dbname=%s port=%s sslmode=%s\n", host, user, dbname, port, sslmode) // Log DB connection details

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbname, port, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to remote database: %v", err)
	}
	fmt.Println("Successfully connected to the remote database")
}

func connectToLocalDB() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	port := os.Getenv("POSTGRES_PORT")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	fmt.Printf("Connecting to local DB with host=%s user=%s dbname=%s port=%s sslmode=%s\n", host, user, dbname, port, sslmode) // Log DB connection details

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host, user, password, dbname, port, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to local database: %v", err)
	}
	fmt.Println("Successfully connected to the local database")
}
