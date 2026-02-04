package main

import (
	"log"

	"github.com/custordev/shortener-api/initializers"
	"github.com/custordev/shortener-api/models"
)

var err * error

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

if err != nil{
		log.Fatal("Failled to run Migrations ")
	}	

	initializers.Db.AutoMigrate(&models.User{},models.Link{},models.User{})
	log.Println("✅ Database migration completed Successfully !")
}