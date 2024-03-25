package main

import (
	"final-project/database"
	"final-project/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetDB()
	router := gin.Default()

	routers.UserRouter(router)
	routers.PhotoRouter(router, db)
	routers.CommentRouter(router, db)
	routers.SocialMediaRouter(router, db)

	router.Run(":5432")
}
