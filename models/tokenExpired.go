package models

import "gorm.io/gorm"

type TokenExpired struct {
	gorm.Model
	Token string `gorm:"type:varchar(255);unique;not null"`
}
