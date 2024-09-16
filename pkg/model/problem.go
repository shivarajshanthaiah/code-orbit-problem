package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model

	Title       string `gorm:"not null"`
	Description string
	Difficulty  string
	Tags        string
}
