package service

import (
	"context"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
)

// InsertTestCasesService implements interfaces.ProblemServiceInter.
func (pr *ProblemService) InsertTestCasesService(ctx context.Context, req *pb.TestCaseRequest) (*pb.ProblemResponse, error) {
	var testCases []model.TestCase
	for _, tc := range req.TestCases {
		testCases = append(testCases, model.TestCase{
			Input:          tc.Input,
			ExpectedOutput: tc.ExpectedOutput,
		})
	}

	testCasesDoc := model.TestCasesDocument{
		ProblemID: int(req.ProblemId),
		TestCases: testCases,
	}

	testCaseID, err := pr.TestCaseRepo.InsertTestCases(ctx, testCasesDoc)
	if err != nil {
		return &pb.ProblemResponse{
			Status:  pb.ProblemResponse_ERROR,
			Message: "Error creating testcases",
			Payload: &pb.ProblemResponse_Error{
				Error: err.Error(),
			},
		}, err
	}

	return &pb.ProblemResponse{
		Status:  pb.ProblemResponse_OK,
		Message: "testcases created succesfully",
		Payload: &pb.ProblemResponse_Data{
			Data: testCaseID, // Return the test case ID in the response
		},
	}, nil
}


func (pr *ProblemService) UpdateTestCasesService(ctx context.Context, req *pb.UpdateTestCaseRequest) (*pb.ProblemResponse, error) {
    var testCases []model.TestCase
    for _, tc := range req.TestCases {
        testCases = append(testCases, model.TestCase{
            Input:          tc.Input,
            ExpectedOutput: tc.ExpectedOutput,
        })
    }

    testCasesDoc := model.TestCasesDocument{
        ProblemID: int(req.ProblemId),
        TestCases: testCases,
    }

    err := pr.TestCaseRepo.UpdateTestCases(ctx, req.TestCaseId, testCasesDoc)
    if err != nil {
        return &pb.ProblemResponse{
            Status:  pb.ProblemResponse_ERROR,
            Message: "Error updating testcases",
            Payload: &pb.ProblemResponse_Error{
                Error: err.Error(),
            },
        }, err
    }

    return &pb.ProblemResponse{
        Status:  pb.ProblemResponse_OK,
        Message: "Test cases updated successfully",
    }, nil
}
