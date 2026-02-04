package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB


func ConnectToDb() {
	var err error
// Build DSN(Datasource name) from environment variables
	dbUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST",),
		os.Getenv("POSTGRES_USER",),
		os.Getenv("POSTGRES_PASSWORD",),
		os.Getenv("POSTGRES_DB", ),
		os.Getenv("POSTGRES_PORT",),
		os.Getenv("POSTGRES_PGSSLMODE",),	
	)
	Db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
 
	if err != nil {
		log.Fatal("Failed to connect to the DB:", err)
	}

	// Test the connection
	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatal("Failed to get DB instance:", err)
	}
 
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
 
	fmt.Println("✅ Successfully connected to database!")
}