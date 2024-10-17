package interfaces

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

type ProblemServiceInter interface {
	InsertProblemService(p *pb.Problem) (*pb.ProblemResponse, error)
	FindAllProblemsService(p *pb.ProbNoParam) (*pb.ProblemList, error)
	EditProblemService(p *pb.Problem) (*pb.Problem, error)
	FindProblemByIDService(p *pb.ProblemId) (*pb.Problem, error)
	UpgradeProblemService(p *pb.ProblemId) (*pb.ProblemResponse, error)

	InsertTestCasesService(ctx context.Context, req *pb.TestCaseRequest) (*pb.ProblemResponse, error)
	UpdateTestCasesService(ctx context.Context, req *pb.UpdateTestCaseRequest) (*pb.ProblemResponse, error)
	GetProblemWithTestCasesService(ctx context.Context, req *pb.ProblemId) (*pb.GetProblemResponse, error)

	SubmitCodeService(ctx context.Context, req *pb.SubmissionRequest) (*pb.SubmissionResponse, error)
	GetUserStats(ctx context.Context, p *pb.UserID) (*pb.StatsResponse, error)
	GetProblemStatsService(req *pb.ProblemStatsRequest) (*pb.ProblemStatsResponse, error)
	GetLeaderboardService() (*pb.LeaderboardResponse, error)
}
