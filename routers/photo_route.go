package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PhotoRouter(route *gin.Engine, db *gorm.DB) {
	PhotoController := &controllers.PhotoController{}
	photo := route.Group("/photos")
	photo.Use(middlewares.Authentication())
	photo.POST("/", PhotoController.CreatePhoto)
	photo.GET("/", PhotoController.GetAllPhotos)
	photo.PUT("/:photoId", middlewares.PhotoAuthorization(), PhotoController.UpdatePhoto)
	photo.DELETE("/:photoId", middlewares.PhotoAuthorization(), PhotoController.DeletePhoto)
}
