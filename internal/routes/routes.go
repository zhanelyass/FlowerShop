package routes

import (
	"FlowerShop/internal/auth"
	"FlowerShop/internal/db"
	"FlowerShop/internal/delivery"
	"FlowerShop/internal/middleware"
	"FlowerShop/internal/repository"
	"FlowerShop/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Открытые маршруты авторизации
	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}

	// Защищённые маршруты
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthRequired())
	{
		// Информация о текущем пользователе
		protected.GET("/me", auth.Me)

		// Инициализация зависимостей цветочного модуля
		flowerRepo := repository.NewFlowerRepository(db.DB)
		flowerService := services.NewFlowerService(flowerRepo)
		flowerHandler := delivery.NewFlowerHandler(flowerService)

		// Доступ для всех авторизованных пользователей
		flowers := protected.Group("/flowers")
		{
			flowers.GET("/", flowerHandler.GetAllFlowers) // Получить все цветы
			flowers.GET("/:id", flowerHandler.GetFlower)  // Получить цветок по ID
		}

		// Только для админа: создание, обновление, удаление цветов
		adminFlowers := protected.Group("/flowers")
		adminFlowers.Use(middleware.RoleRequired("admin"))
		{
			adminFlowers.POST("/", flowerHandler.CreateFlower)      // Создать цветок
			adminFlowers.PUT("/:id", flowerHandler.UpdateFlower)    // Обновить цветок
			adminFlowers.DELETE("/:id", flowerHandler.DeleteFlower) // Удалить цветок
		}
	}
}
