package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model

	Title       string `gorm:"not null;size:255"`
	Description string `gorm:"type:text"`
	Difficulty  string
	Type        string
	IsPremium   bool `gorm:"default:false"`
}
