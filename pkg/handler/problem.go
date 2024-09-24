package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

func (pr *ProblemHandler) InsertProblem(ctx context.Context, p *pb.Problem) (*pb.ProblemResponse, error) {
	reponse, err := pr.SVC.InsertProblemService(p)
	if err != nil {
		return reponse, err
	}
	return reponse, nil
}

func (pr *ProblemHandler) GetAllProblems(ctx context.Context, p *pb.ProbNoParam) (*pb.ProblemList, error) {
	response, err := pr.SVC.FindAllProblemsService(p)
	if err != nil {
		return response, nil
	}
	return response, nil
}

func (ph *ProblemHandler) GetProblemWithTestCases(ctx context.Context, p *pb.ProblemId) (*pb.GetProblemResponse, error) {
	response, err := ph.SVC.GetProblemWithTestCasesService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ph *ProblemHandler) FindProblemByID(ctx context.Context, p *pb.ProblemId) (*pb.Problem, error) {
	response, err := ph.SVC.FindProblemByIDService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ph *ProblemHandler) EditProblem(ctx context.Context, p *pb.Problem) (*pb.Problem, error) {
	response, err := ph.SVC.EditProblemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ph *ProblemHandler) AdminUpgradeProbem(ctx context.Context, p *pb.ProblemId) (*pb.ProblemResponse, error) {
	response, err := ph.SVC.UpgradeProblemService(p)
	if err != nil {
		return response, err
	}
	return response, nil
}
