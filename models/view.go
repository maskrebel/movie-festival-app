package models

import "gorm.io/gorm"

type View struct {
	gorm.Model
	UserID   *uint `gorm:"index"`
	MovieID  uint  `gorm:"not null;index"`
	Duration int   `gorm:"not null"`
}
