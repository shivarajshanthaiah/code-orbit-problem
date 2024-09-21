package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

func (ph *ProblemHandler) InsertTestCases(ctx context.Context, p *pb.TestCaseRequest) (*pb.ProblemResponse, error) {
	response, err := ph.SVC.InsertTestCasesService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ph *ProblemHandler) UpdateTestCases(ctx context.Context, p *pb.UpdateTestCaseRequest) (*pb.ProblemResponse, error){
	response, err := ph.SVC.UpdateTestCasesService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}
