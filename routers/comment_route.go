package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRouter(route *gin.Engine, db *gorm.DB) {
	CommentController := &controllers.CommentController{}
	comment := route.Group("/comments")
	comment.Use(middlewares.Authentication())
	comment.POST("/", CommentController.CreateComment)
	comment.PUT("/:commentId", middlewares.CommentAuthorization(), CommentController.UpdateComment)
	comment.GET("/", CommentController.GetAllComments)
	comment.DELETE("/:commentId", middlewares.CommentAuthorization(), CommentController.DeleteComment)
}
