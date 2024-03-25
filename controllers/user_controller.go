package controllers

import (
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) UserLogin(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)

	user := models.User{}
	password := ""

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	password = user.Password

	err := uc.db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"data": gin.H{
				"error":   "unauthorized",
				"message": "invalid email/password",
			},
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"data": gin.H{
				"error":   "unauthorized",
				"message": "invalid email/password",
			},
		})
		return
	}
	token := helpers.GenerateToken(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"token": token,
		},
	})

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	user := models.User{}
	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	temp, _ := strconv.Atoi(ctx.Param("userId"))
	user.ID = int(temp)
	user.Updated_at = time.Now()

	fmt.Printf("Value Update: %+v\n", user)
	err := uc.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Update User Data",
			},
		})
		return
	}
	updatedUser := models.User{}
	_ = uc.db.First(&updatedUser, "id = ?", user.ID).Error

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":         updatedUser.ID,
			"email":      updatedUser.Email,
			"username":   updatedUser.Username,
			"age":        updatedUser.Age,
			"updated_at": updatedUser.Updated_at,
		},
	})
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	user := models.User{}

	contentType := helpers.GetContentType(ctx)

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	err := uc.db.Debug().Create(&user).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Create User",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"age":      user.Age,
			"email":    user.Email,
			"id":       user.ID,
			"username": user.Username,
		},
	})

}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	UserData := ctx.MustGet("userData").(jwt.MapClaims)
	userIdFromJwt := UserData["id"].(float64)
	fmt.Println(userIdFromJwt)
	user := models.User{}

	err := uc.db.Where("id= ?", userIdFromJwt).Delete(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Delete User",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your Account Has Successfully Deleted",
		},
	})
}
