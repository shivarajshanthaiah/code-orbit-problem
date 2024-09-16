package repo

import (
	inter "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/repo/interfaces"
	"gorm.io/gorm"
)

type ProblemRepository struct {
	DB *gorm.DB
}

func NewProblemRepository(db *gorm.DB) inter.ProblemRepoInter {
	return &ProblemRepository{
		DB: db,
	}
}
