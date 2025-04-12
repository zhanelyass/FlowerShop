package routes

import (
	"FlowerShop/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	flowers := r.Group("api/v1/auth")
	{
		flowers.POST("/login", auth.login)
		flowers.POST("/register", auth.Register)
	}
}
