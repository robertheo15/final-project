package entity

type Comment struct {
	GormModel
	UserId  int    `gorm:"not null" json:"user_id"`
	PhotoId int    `gorm:"not null" json:"photo_id"`
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	User    *User
	Photo   *Photo
}
