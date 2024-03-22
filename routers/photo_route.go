package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(route *gin.Engine) {
	photo := route.Group("/photos")
	photo.Use(middlewares.Authentication())
	photo.POST("/", controllers.CreatePhoto)
	photo.GET("/", controllers.GetAllPhotos)
}
