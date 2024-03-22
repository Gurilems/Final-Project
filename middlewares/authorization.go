package middlewares

import (
	"final-project/helpers"
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
