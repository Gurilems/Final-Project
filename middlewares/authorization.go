package middlewares

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data": gin.H{
					"error": err.Error(),
					"msg":   "unauthenticated",
				},
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}

func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserId, err := strconv.Atoi(c.Param("userId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": err.Error(),
					"msg":   "Data not found",
				},
			})
			return
		}

		UserData := c.MustGet("userData").(jwt.MapClaims)
		userIdFromJwt := UserData["id"].(float64)
		fmt.Println("ini user jwt : ", userIdFromJwt)
		fmt.Println("ini user param : ", UserId)

		if UserId != int(userIdFromJwt) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Unauthorized",
					"msg":   "You are Not Allowed to Access This Data",
				},
			})
			return
		}
		c.Next()
	}
}

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := database.GetDB()
		photo := models.Photo{}
		photoId, err := strconv.Atoi(c.Param("photoId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": err.Error(),
					"msg":   "Data does'nt exist",
				},
			})
			return
		}

		UserData := c.MustGet("userData").(jwt.MapClaims)
		userIdFromJwt := UserData["id"].(float64)

		err = db.Select("UserID").First(&photo, uint(photoId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": err.Error(),
					"msg":   "Data does'nt exist",
				},
			})
			return
		}

		fmt.Println(photo.UserID, userIdFromJwt)

		if photo.UserID != uint(userIdFromJwt) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Unauthorized",
					"msg":   "You are Not Allowed to Access This Data",
				},
			})
			return
		}
		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := database.GetDB()
		comment := models.Comment{}
		commentId, err := strconv.Atoi(c.Param("commentId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Data not found",
					"msg":   "Data does'nt exist",
				},
			})
			return
		}

		UserData := c.MustGet("userData").(jwt.MapClaims)
		userIdFromJwt := UserData["id"].(float64)

		err = db.Select("UserID").First(&comment, uint(commentId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Data not found",
					"msg":   "Data does'nt exist",
				},
			})
			return
		}

		if comment.UserID != uint(userIdFromJwt) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Unauthorized",
					"msg":   "You are Not Allowed to Access This Data",
				},
			})
			return
		}
		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := database.GetDB()
		socialMedia := models.SocialMedia{}
		socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Data not found",
					"msg":   "Data does'nt exist",
				},
			})
			return
		}

		UserData := c.MustGet("userData").(jwt.MapClaims)
		userIdFromJwt := UserData["id"].(float64)

		err = db.Select("UserID").First(&socialMedia, uint(socialMediaId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Data not found",
					"msg":   "Data does'nt exist",
				},
			})
			return
		}

		if socialMedia.UserID != uint(userIdFromJwt) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": gin.H{
					"error": "Unauthorized",
					"msg":   "You are Not Allowed to Access This Data",
				},
			})
			return
		}
		c.Next()
	}
}
