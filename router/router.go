package router

import (
	"GinTutorial/controllers"
	"GinTutorial/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/login", controllers.Login)
				auth.POST("/register", controllers.Register)
			}

			ex := v1.Group("/ex")
			{
				ex.GET("/exchangeRates", controllers.GetExchangeRates)
				ex.Use(middlewares.AuthMiddleWare())
				ex.POST("/exchangeRates", controllers.CreateExchangeRate)
			}
		}
	}
	return r
}
