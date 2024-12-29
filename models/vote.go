package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserID  uint `gorm:"not null;index:idx_user_movie_id,unique" json:"user_id"`
	MovieID uint `gorm:"not null;index:idx_user_movie_id,unique" json:"movie_id"`
}
