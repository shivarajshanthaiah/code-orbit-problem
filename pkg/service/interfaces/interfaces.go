package interfaces

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

type ProblemServiceInter interface {
	InsertProblemService(p *pb.Problem) (*pb.ProblemResponse, error)
	FindAllProblemsService(p *pb.ProbNoParam) (*pb.ProblemList, error)

	InsertTestCasesService(ctx context.Context, req *pb.TestCaseRequest) (*pb.ProblemResponse, error)
	UpdateTestCasesService(ctx context.Context, req *pb.UpdateTestCaseRequest) (*pb.ProblemResponse, error)
}
