package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"movie-festival-app/utils"
	"reflect"
)

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

func (m Movie) MarshalJSON() ([]byte, error) {
	type Alias Movie
	alias := Alias(m)
	value := reflect.ValueOf(alias)
	return json.Marshal(utils.ConvertToJson(value))
}
