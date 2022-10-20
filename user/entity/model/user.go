package model

import (
	"final-project/comment/entity/model"
	"final-project/common/base/entity/models"
	"final-project/common/helpers"
	model2 "final-project/photo/entity/model"
	model3 "final-project/social/entity/model"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	models.GormModel
	Username string          `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string          `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required, email~Invalid email format,email~Invalid format email"`
	Password string          `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int             `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,numeric~Fill age with number,range(8|99)~minimum 8 years old"`
	Photos   []model2.Photo  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments []model.Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Socials  []model3.Social `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socials"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)
	if errCreate != nil {
		err = errCreate
		return err
	}
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return
}
