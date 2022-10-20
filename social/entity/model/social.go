package model

import (
	"final-project/common/base/entity/models"
	"final-project/user/entity/model"
)

type Social struct {
	models.GormModel
	Name   string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	Url    string `gorm:"not null" json:"url" form:"url" valid:"required~Url is required"`
	UserId int    `gorm:"not null" json:"user_id"`
	User   *model.User
}
