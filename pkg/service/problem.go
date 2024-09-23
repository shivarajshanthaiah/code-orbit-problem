package service

import (
	"context"

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
		IsPremium:   p.IsPremium,
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
			IsPremium:   problem.IsPremium,
		}
		problemList.Problems = append(problemList.Problems, pbProblem)
	}

	return &problemList, nil
}

func (pr *ProblemService) GetProblemWithTestCasesService(ctx context.Context, req *pb.ProblemId) (*pb.GetProblemResponse, error) {
	// Fetch the problem from PostgreSQL
	problem, err := pr.Repo.GetProblemByID(req.ID)
	if err != nil {
		return &pb.GetProblemResponse{
			Status:  pb.GetProblemResponse_ERROR,
			Message: "Error fetching problem",
			Payload: &pb.GetProblemResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Fetch the test cases from MongoDB
	testCasesDoc, err := pr.TestCaseRepo.GetTestCasesByProblemID(req.ID)
	if err != nil {
		return &pb.GetProblemResponse{
			Status:  pb.GetProblemResponse_ERROR,
			Message: "Error fetching test cases",
			Payload: &pb.GetProblemResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	// Convert the problem and test cases into the response format
	var grpcTestCases []*pb.TestCase
	for _, tc := range testCasesDoc.TestCases { // Now this should work
		grpcTestCases = append(grpcTestCases, &pb.TestCase{
			TestCaseId:     tc.ID, // Assuming this is the field for MongoDB ObjectID
			Input:          tc.Input,
			ExpectedOutput: tc.ExpectedOutput,
		})
	}

	return &pb.GetProblemResponse{
		Status:  pb.GetProblemResponse_OK,
		Message: "Problem and test cases fetched successfully",
		Payload: &pb.GetProblemResponse_Data{
			Data: &pb.ProblemWithTestCases{
				Problem: &pb.Problem{
					ID:          uint32(problem.ID),
					Title:       problem.Title,
					Discription: problem.Description,
					Difficulty:  problem.Difficulty,
					Tags:        problem.Tags,
					IsPremium:   problem.IsPremium,
				},
				TestCases: grpcTestCases,
			},
		},
	}, nil
}

func (pr *ProblemService) EditProblemService(p *pb.Problem) (*pb.Problem, error) {
	problem, err := pr.Repo.FindProblemByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	problem.Title = p.Title
	problem.Description = p.Discription
	problem.Difficulty = p.Difficulty
	problem.Tags = p.Tags
	problem.IsPremium = p.IsPremium

	err = pr.Repo.UpdateProblem(problem)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pr *ProblemService) FindProblemByIDService(p *pb.ProblemId) (*pb.Problem, error) {
	problem, err := pr.Repo.FindProblemByID(uint(p.ID))
	if err != nil {
		return nil, err
	}

	return &pb.Problem{
		ID:          uint32(problem.ID),
		Title:       problem.Title,
		Discription: problem.Description,
		Difficulty:  problem.Difficulty,
		Tags:        problem.Difficulty,
		IsPremium:   problem.IsPremium,
	}, nil
}