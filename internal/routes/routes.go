package routes

import (
	"FlowerShop/internal/delivery"
	"FlowerShop/internal/repository"
	"FlowerShop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Repository initialization
	flowerRepo := repository.NewFlowerRepository(db)

	// Service initialization
	flowerService := services.NewFlowerService(flowerRepo)

	// Handler initialization
	flowerHandler := delivery.NewFlowerHandler(flowerService)

	flowers := r.Group("api/v1/flowers")
	{
		flowers.GET("/", flowerHandler.GetAllFlowers)
		flowers.GET("/:id", flowerHandler.GetFlower)
		flowers.POST("/", flowerHandler.CreateFlower)
		flowers.PUT("/:id", flowerHandler.UpdateFlower)
		flowers.DELETE("/:id", flowerHandler.DeleteFlower)
	}
}
