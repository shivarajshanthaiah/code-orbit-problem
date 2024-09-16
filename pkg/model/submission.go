package model

import "gorm.io/gorm"

type Submission struct {
	gorm.Model

	UserID       string `gorm:"not null"`
	ProblemID    int    `gorm:"not null"`
	Language     string
	Code         string
	Status       string
	AttemptCount int
}
