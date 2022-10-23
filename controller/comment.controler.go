package controller

import "gorm.io/gorm"

type CommentController struct {
	db *gorm.DB
}

func NewPCommentController(db *gorm.DB) *CommentController {
	return &CommentController{
		db: db,
	}
}
