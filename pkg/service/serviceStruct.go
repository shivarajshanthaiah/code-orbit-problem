package service

import (
	interRepo "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/service/interfaces"
)

type ProblemService struct {
	Repo interRepo.ProblemRepoInter
}

func NewProblemService(repo interRepo.ProblemRepoInter) interfaces.ProblemServiceInter {
	return &ProblemService{
		Repo: repo,
	}
}
