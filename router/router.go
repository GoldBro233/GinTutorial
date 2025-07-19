package router

import (
	"GinTutorial/controllers"

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
		}
	}
	return r
}
