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

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db: db}
}

func (pc *PhotoController) CreatePhoto(ctx *gin.Context) {
	photo := models.Photo{}

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataID := uint(userData["id"].(float64))

	photo.Created_at = time.Now()
	photo.Updated_at = time.Now()
	photo.UserID = userDataID

	err := pc.db.Create(&photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Create Photo",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.Photo_url,
			"user_id":    photo.UserID,
			"created_at": photo.Created_at,
		},
	})
}

func (pc *PhotoController) GetAllPhotos(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataID := uint(userData["id"].(float64))

	photos := []models.Photo{}
	err := pc.db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("ID", "email", "username", "created_at", "updated_at")
	}).Where("user_id = ?", userDataID).Find(&photos).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Get Photo List",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   photos,
	})
}

func (pc *PhotoController) UpdatePhoto(ctx *gin.Context) {
	photo := models.Photo{}
	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	temp, _ := strconv.Atoi(ctx.Param("photoId"))
	photo.ID = uint(temp)
	photo.Updated_at = time.Now()

	fmt.Printf("Value Update: %+v\n", photo)
	err := pc.db.Model(&photo).Where("id = ?", photo.ID).Updates(&photo).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Update Photo",
			},
		})
		return
	}

	updatedPhoto := models.Photo{}
	_ = pc.db.First(&updatedPhoto, "id = ?", photo.ID).Error

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":         updatedPhoto.ID,
			"title":      updatedPhoto.Title,
			"caption":    updatedPhoto.Caption,
			"photo_url":  updatedPhoto.Photo_url,
			"user_id":    updatedPhoto.UserID,
			"updated_at": updatedPhoto.Updated_at,
		},
	})
}

func (pc *PhotoController) DeletePhoto(ctx *gin.Context) {
	photo := models.Photo{}
	temp, _ := strconv.Atoi(ctx.Param("photoId"))

	photo.ID = uint(temp)

	err := pc.db.Where("id= ?", photo.ID).Delete(&photo).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Delete Photo",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your Photo Has been Successfully Deleted",
		},
	})
}
