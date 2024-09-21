package service

import (
	interRepo "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/service/interfaces"
)

type ProblemService struct {
	Repo         interRepo.ProblemRepoInter
	TestCaseRepo interRepo.MongoRepoInter
}

func NewProblemService(repo interRepo.ProblemRepoInter, testRepo interRepo.MongoRepoInter) interfaces.ProblemServiceInter {
	return &ProblemService{
		Repo:         repo,
		TestCaseRepo: testRepo,
	}
}
