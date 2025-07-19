package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/login", func(context *gin.Context) {
					context.JSON(200, gin.H{
						"msg": "Login successfully",
					})
				})
				auth.POST("/register")
			}
		}
	}
	return r
}
