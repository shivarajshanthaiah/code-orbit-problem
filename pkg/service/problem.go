package service

import (
	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

// InsertProblemService implements interfaces.ProblemServiceInter.
func (pr *ProblemService) InsertProblemService(p *pb.Problem) (*pb.ProblemResponse, error) {
	problem := model.Problem{
		Title:       p.Title,
		Description: p.Discription,
		Difficulty:  p.Difficulty,
		Tags:        p.Tags,
	}

	err := pr.Repo.InsertProblem(&problem)
	if err != nil {
		return &pb.ProblemResponse{
			Status:  pb.ProblemResponse_ERROR,
			Message: "Error creating problem",
			Payload: &pb.ProblemResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.ProblemResponse{
		Status:  pb.ProblemResponse_OK,
		Message: "Problem created succesfully",
	}, nil
}
