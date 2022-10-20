package model

import (
	"final-project/common/base/entity/models"
	"final-project/photo/entity/model"
	model2 "final-project/user/entity/model"
)

type Comment struct {
	models.GormModel
	UserId  int    `gorm:"not null" json:"user_id"`
	PhotoId int    `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	User    *model2.User
	Photo   *model.Photo
}
