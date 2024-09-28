package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

func (ph *ProblemHandler) SubmitCode(ctx context.Context, p *pb.SubmissionRequest) (*pb.SubmissionResponse, error) {
	response, err := ph.SVC.SubmitCodeService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ph *ProblemHandler) GetUserStats(ctx context.Context, p *pb.UserID) (*pb.StatsResponse, error) {
	response, err := ph.SVC.GetUserStats(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}
