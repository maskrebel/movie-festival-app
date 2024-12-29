package models

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	UserID  uint `gorm:"not null" json:"user_id"`
	MovieID uint `gorm:"not null" json:"movie_id"`
}
