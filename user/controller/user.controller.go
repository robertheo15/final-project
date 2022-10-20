package controller

import (
	helpers2 "final-project/common/helpers"
	"final-project/user/entity/model"
	"final-project/user/entity/response"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	user := model.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helpers2.BadRequestResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		helpers2.BadRequestResponse(ctx, err.Error())
		return
	}

	err = controller.db.Create(&user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helpers2.NotFoundResponse(ctx, err.Error())
			return
		}
		helpers2.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helpers2.WriteJsonResponse(ctx, http.StatusCreated, response.UserCreateResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	})
}

func (controller *UserController) UserLogin(ctx *gin.Context) {
	user := model.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helpers2.BadRequestResponse(ctx, err.Error())
		return
	}

	password := user.Password
	err = controller.db.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		helpers2.WriteJsonResponse(ctx, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "username / password is not match",
		})
		return
	}

	comparePass := helpers2.ComparePassword(user.Password, password)
	if !comparePass {
		helpers2.WriteJsonResponse(ctx, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "username / password is not match",
		})
		return
	}
	token := helpers2.GenerateToken(user.ID, user.Email)
	ctx.JSON(http.StatusOK, response.UserLoginResponse{
		Token: token,
	})
}
