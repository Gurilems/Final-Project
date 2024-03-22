package routers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataID := uint(userData["id"].(float64))

	comment.Created_at = time.Now()
	comment.Updated_at = time.Now()
	comment.UserID = userDataID

	if err := db.Create(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Create Comment",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.Created_at,
		},
	})
}

func GetAllComments(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataID := uint(userData["id"].(float64))
	comments := []models.Comment{}

	err := db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("ID", "email", "username", "created_at", "updated_at")
	}).Preload("Photo", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("ID", "title", "caption", "photo_url", "user_id", "created_at", "updated_at")
	}).Where("user_id = ?", userDataID).Find(&comments).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Get Comment List",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments,
	})
}

func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Invalid comment ID",
			},
		})
		return
	}
	comment.ID = uint(commentID)
	comment.Updated_at = time.Now()

	if err := db.Model(&comment).Where("id = ?", comment.ID).Updates(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Update Comment",
			},
		})
		return
	}

	updatedComment := models.Comment{}
	if err := db.First(&updatedComment, "id = ?", comment.ID).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to fetch updated comment",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   updatedComment,
	})
}

func DeleteComment(ctx *gin.Context) {
	db := database.GetDB()
	comment := models.Comment{}
	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Invalid comment ID",
			},
		})
		return
	}

	comment.ID = uint(commentID)

	if err := db.Where("ID= ?", comment.ID).Delete(&comment).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Delete Comment",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your Comment Has been Successfully Deleted",
		},
	})
}
