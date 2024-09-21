package interfaces

import (
	"context"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
)

type ProblemRepoInter interface {
	InsertProblem(Problem *model.Problem) error
	FindProblemByID(ProblemID uint) (*model.Problem, error)
	GetAllProblems() (*[]model.Problem, error)
	GetProblemByID(problemID int32) (*model.Problem, error)
}

type MongoRepoInter interface {
	InsertTestCases(ctx context.Context, testCasesDoc model.TestCasesDocument) (string, error)
	UpdateTestCases(ctx context.Context, testCaseID string, testCasesDoc model.TestCasesDocument) error
	GetTestCasesByProblemID(problemID int32) (*model.TestCasesDocument, error)
}
