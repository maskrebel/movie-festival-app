package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"not null;unique"`
	Year        int    `gorm:"not null"`
	Description string `gorm:"not null"`
	Duration    int    `gorm:"not null"`
	Artist      string `gorm:"not null"`
	Genre       string `gorm:"not null"`
	Url         string `gorm:"not null;unique"`
	Views       int    `gorm:"default:0"`
}
