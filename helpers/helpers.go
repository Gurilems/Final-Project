package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHashing(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hash)
}

func ComparePass(h, p []byte) bool {
	err := bcrypt.CompareHashAndPassword(h, p)
	return err == nil
}

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}

func GenerateToken(userId int, email string) string {
	claims := jwt.MapClaims{
		"id":    userId,
		"email": email,
		"time":  time.Now().String(),
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte("rahasia"))
	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte("kepodeh"), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil

}
