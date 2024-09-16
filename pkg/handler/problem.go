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