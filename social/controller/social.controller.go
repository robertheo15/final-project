package controller

import (
	"final-project/common/helpers"
	"final-project/social/entity/model"
	"final-project/social/entity/response"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type SocialController struct {
	db *gorm.DB
}

func NewSocialController(db *gorm.DB) *SocialController {
	return &SocialController{
		db: db,
	}
}

func (controller *SocialController) CreateSocial(ctx *gin.Context) {
	social := model.Social{}

	err := ctx.ShouldBindJSON(&social)
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&social)
	if err != nil {
		helpers.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Create(&social).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers.NotFoundResponse(ctx, err.Error())
			return
		}
		helpers.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helpers.WriteJsonResponse(ctx, http.StatusCreated, response.SocialCreateResponse{
		Id:             social.ID,
		Name:           social.Name,
		SocialMediaUrl: social.Url,
	})
}
