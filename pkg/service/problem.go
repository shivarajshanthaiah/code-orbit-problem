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
		IsPremium:   p.Is_Premium,
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

func (pr *ProblemService) FindAllProblemsService(p *pb.ProbNoParam) (*pb.ProblemList, error) {
	result, err := pr.Repo.GetAllProblems()
	if err != nil {
		return nil, err
	}

	// Check if result is nil
	if result == nil {
		return &pb.ProblemList{}, nil
	}

	var problemList pb.ProblemList
	for _, problem := range *result {
		pbProblem := &pb.Problem{
			ID:          uint32(problem.ID),
			Title:       problem.Title,
			Discription: problem.Description,
			Difficulty:  problem.Difficulty,
			Tags:        problem.Tags,
			Is_Premium:  problem.IsPremium,
		}
		problemList.Problems = append(problemList.Problems, pbProblem)
	}

	return &problemList, nil
}
