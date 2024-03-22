package main

import (
	"final-project/database"
	"final-project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()
	router := gin.Default()

	routers.UserRouter(router)
	routers.PhotoRouter(router)
	// routers.CommentRouter(router)
	routers.SocialMediaRouter(router)

	router.Run(":8080")
}
