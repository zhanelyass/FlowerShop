package main

import (
	"FlowerShop/internal/models"
	"FlowerShop/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://myuser:mypassword@localhost:5444/flowerdatabase?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.AutoMigrate(&models.Flower{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run(":8081")

}
