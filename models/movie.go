package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string `gorm:"not null;unique"`
	Year        int    `gorm:"not null"`
	Description string `gorm:"not null"`
	Duration    int    `gorm:"not null"`
	Artists     string `gorm:"not null"`
	Genres      string `gorm:"not null"`
	WatchURL    string `gorm:"not null;unique"`
	Views       int    `gorm:"default:0"`
	Votes       int    `gorm:"default:0"`
}
