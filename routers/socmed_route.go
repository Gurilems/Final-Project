package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func SocialMediaRouter(route *gin.Engine) {
	sosmed := route.Group("/socialmedias")
	sosmed.Use(middlewares.Authentication())
	sosmed.GET("/", controllers.GetAllSocialMedia)
	sosmed.POST("/", controllers.CreateSocialMedia)
}
