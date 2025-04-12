package main

import (
	"FlowerShop/internal/db"
	"FlowerShop/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8081")

}
