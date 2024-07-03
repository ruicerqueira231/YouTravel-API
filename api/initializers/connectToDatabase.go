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
		env = "production"
		fmt.Println(env)
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
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"),
	)
}

func buildRemoteDSN() string {
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "ep-small-fire-a202x97i-pooler.eu-central-1.aws.neon.tech"
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		host = "default"
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
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"),
	)
}
