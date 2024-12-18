package interfaces

import (
	"context"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
)

type ProblemRepoInter interface {
	InsertProblem(Problem *model.Problem) error
	FindProblemByID(ProblemID uint) (*model.Problem, error)
	GetAllProblems() (*[]model.Problem, error)
	GetProblemByID(problemID uint) (*model.Problem, error)
	UpdateProblem(problem *model.Problem) error

	SaveSubmission(submission *model.Submission) error
	FetchSubmission(userID string, problemID int) (*model.Submission, error)
	UpdateSubmission(submission *model.Submission) error
	GetUserStats(userID string) (map[string]int, error)

	CountByCondition(condition string) int64
	CountByDifficulty(difficulty string) int64
	GetDistinctProblemTypes() ([]string, error)
	CountByPremiumStatus(isPremium bool) int64
	GetLeaderboard() ([]model.LeaderboardEntry, error)
}

type MongoRepoInter interface {
	InsertTestCases(ctx context.Context, testCasesDoc model.TestCasesDocument) (string, error)
	UpdateTestCases(ctx context.Context, testCaseID string, testCasesDoc model.TestCasesDocument) error
	GetTestCasesByProblemID(problemID int32) (*model.TestCasesDocument, error)
}
