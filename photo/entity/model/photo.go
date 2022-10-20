package model

import (
	"final-project/comment/entity/model"
	"final-project/common/base/entity/models"
	model2 "final-project/user/entity/model"
)

type Photo struct {
	models.GormModel
	Title    string          `gorm:"not null" json:"username" form:"username" valid:"required~Title is required"`
	Caption  string          `gorm:"not null" json:"email" form:"email" valid:"required~Caption is required"`
	PhotoUrl string          `gorm:"not null" json:"password" form:"password" valid:"required~Photo url is required"`
	UserId   int             `gorm:"not null" json:"user_id"`
	Comment  []model.Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	User     *model2.User
}
