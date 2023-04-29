package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// use os package to get the env variable which is already set
func envVariable(key string) string {

	// set env variable using os package
	os.Setenv(key, "gopher")

	// return the env variable using os package
	return os.Getenv(key)
}

// var DSN = "host=localhost user=felixvnolasco password=holamundo dbname=state port=5432 sslmode=disable"

var DB *gorm.DB

func DBConnection() {

	// value := envVariable(DSN)

	// Load environment variables from file.

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	var error error
	// DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	DB, error = gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if error != nil {
		// log.Fatal(error)
		log.Fatalf("failed to connect to PlanetScale: %v", error)
	} else {
		log.Print("DB Connected")
	}
}
