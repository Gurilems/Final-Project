package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SocialMediaRouter(route *gin.Engine, db *gorm.DB) {
	SocialMediaController := &controllers.SocialMediaController{}
	sosmed := route.Group("/socialmedias")
	sosmed.Use(middlewares.Authentication())
	sosmed.GET("/", SocialMediaController.GetAllSocialMedia)
	sosmed.POST("/", SocialMediaController.CreateSocialMedia)
	sosmed.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), SocialMediaController.UpdateSocialMedia)
	sosmed.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), SocialMediaController.DeleteSocialMedia)
}
