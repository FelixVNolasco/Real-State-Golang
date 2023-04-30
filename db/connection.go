package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DSN = "host=localhost user=felixvnolasco password=holamundo dbname=state port=5432 sslmode=disable"

var DB *gorm.DB

func DBConnection() *gorm.DB {

	// Load environment variables from file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	var err error
	// DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	DB, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		// log.Fatal(error)
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	} else {
		log.Print("DB Connected")
	}

	return DB
}
