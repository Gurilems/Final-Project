package routers

import (
	"final-project/controllers"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.Engine) {
	UserController := &controllers.UserController{}
	user := route.Group("/users")
	user.POST("/register", UserController.CreateUser)
	user.POST("/login", UserController.UserLogin)

	user.Use(middlewares.Authentication())
	user.DELETE("/", UserController.DeleteUser)
	user.PUT("/:userId", middlewares.UserAuthorization(), UserController.UpdateUser)
}
