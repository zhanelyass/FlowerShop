package routes

import (
	"FlowerShop/internal/auth"
	"FlowerShop/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	authRotes := r.Group("api/v1/auth")
	{
		authRotes.POST("/login", auth.login)
		authRotes.POST("/register", auth.Register)
	}

	protected := r.Group("api/v1")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/me", auth.Me)
	}
}
