package controllers

import (
	"final-project/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaController struct {
	db *gorm.DB
}

func NewSocialMediaController(db *gorm.DB) *SocialMediaController {
	return &SocialMediaController{db: db}
}

func (smc *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	sosmed := models.SocialMedia{}
	if err := ctx.ShouldBindJSON(&sosmed); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataId := uint(userData["id"].(float64))

	sosmed.Created_at = time.Now()
	sosmed.Updated_at = time.Now()
	sosmed.UserID = userDataId

	err := smc.db.Create(&sosmed).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Create Social Media",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":               sosmed.ID,
			"name":             sosmed.Name,
			"social_media_url": sosmed.Social_Media_Url,
			"user_id":          sosmed.UserID,
			"created_at":       sosmed.Created_at,
		},
	})
}

func (smc *SocialMediaController) GetAllSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataId := uint(userData["id"].(float64))
	socialMedias := []models.SocialMedia{}
	err := smc.db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("ID", "username", "created_at", "updated_at")
	}).Where("user_id = ?", userDataId).Find(&socialMedias).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Get Social Media List",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   socialMedias,
	})
}

func (smc *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}
	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	temp, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	socialMedia.ID = uint(temp)
	socialMedia.Updated_at = time.Now()

	fmt.Printf("Value Update: %+v\n", socialMedia)
	err := smc.db.Model(&socialMedia).Where("id = ?", socialMedia.ID).Updates(&socialMedia).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Update Social Media",
			},
		})
		return
	}

	updatedSosmed := models.SocialMedia{}
	_ = smc.db.First(&updatedSosmed, "id = ?", socialMedia.ID).Error

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":               updatedSosmed.ID,
			"name":             updatedSosmed.Name,
			"social_media_url": updatedSosmed.Social_Media_Url,
			"user_id":          updatedSosmed.UserID,
			"updated_at":       updatedSosmed.Updated_at,
		},
	})
}

func (smc *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	sosmed := models.SocialMedia{}
	temp, _ := strconv.Atoi(ctx.Param("socialMediaId"))
	sosmed.ID = uint(temp)

	err := smc.db.Where("ID= ?", sosmed.ID).Delete(&sosmed).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": gin.H{
				"error": err.Error(),
				"msg":   "Failed to Delete Social Media",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your Social Media Has been Successfully Deleted",
		},
	})
}
