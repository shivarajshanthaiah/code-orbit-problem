package interfaces

import "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"

type ProblemRepoInter interface {
	InsertProblem(Problem *model.Problem) error
	FindProblemByID(ProblemID uint) (*model.Problem, error)
}
