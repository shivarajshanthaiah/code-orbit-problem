package service

import (
	"context"
	"log"
	"strings"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/model"
	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
	usercodeexcecution "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/userCodeExcecution"
)

func (pr *ProblemService) SubmitCodeService(ctx context.Context, req *pb.SubmissionRequest) (*pb.SubmissionResponse, error) {
	// Step 1: Fetch problem details
	problem, err := pr.Repo.GetProblemByID(uint(req.ProblemId))
	if err != nil {
		return &pb.SubmissionResponse{
			Status:  pb.SubmissionResponse_ERROR,
			Message: "Failed to fetch problem details",
		}, err
	}

	// Step 2: Fetch test cases
	testCasesDoc, err := pr.TestCaseRepo.GetTestCasesByProblemID(req.ProblemId)
	if err != nil {
		return &pb.SubmissionResponse{
			Status:  pb.SubmissionResponse_ERROR,
			Message: "Failed to fetch test cases",
		}, err
	}

	log.Printf("Received code for execution: %s", req.Code)

	// Step 3: Initialize counts for passed and failed test cases
	passedCount := 0
	failedCount := 0

	for _, testCase := range testCasesDoc.TestCases {
		log.Printf("Executing test case with input: %s, Expected output: %s", testCase.Input, testCase.ExpectedOutput)

		// Execute the user's code
		output, execErr := usercodeexcecution.RunUserCode(problem.Type, req.Code, testCase.Input)
		if execErr != nil {
			log.Printf("Execution failed for input: %s, Error: %v", testCase.Input, execErr)
			failedCount++
			continue
		}

		// Log actual output
		log.Printf("Actual output: %s", output)

		// Compare the trimmed output with the expected output
		if strings.TrimSpace(output) == strings.TrimSpace(testCase.ExpectedOutput) {
			passedCount++
			log.Printf("Test case passed")
		} else {
			failedCount++
			log.Printf("Test case failed: input: %s, expected: %s, actual: %s", testCase.Input, testCase.ExpectedOutput, output)
		}
	}

	// Step 5: Determine if the submission passed or failed
	status := "passed"
	message := "Submission passed"
	if failedCount > 0 {
		status = "failed"
		message = "Submission failed"
	}

	// Step 6: Check for existing submission
	existingSubmission, err := pr.Repo.FetchSubmission(req.UserId, int(req.ProblemId))

	var submission *model.Submission
	if err == nil {
		// Update the existing submission with new code and status
		submission = existingSubmission
		submission.Code = req.Code // Update the code
		submission.Status = status // Update the status
		submission.AttemptCount++  // Increment the attempt count
		err = pr.Repo.UpdateSubmission(submission)
		if err != nil {
			return &pb.SubmissionResponse{
				Status:  pb.SubmissionResponse_ERROR,
				Message: "Failed to update submission",
			}, err
		}
	} else {
		// Create a new submission
		submission = &model.Submission{
			UserID:       req.UserId,
			ProblemID:    int(req.ProblemId),
			Language:     req.Language,
			Code:         req.Code,
			Status:       status,
			AttemptCount: 1, // First attempt
		}
		// Save the new submission
		err = pr.Repo.SaveSubmission(submission)
		if err != nil {
			return &pb.SubmissionResponse{
				Status:  pb.SubmissionResponse_ERROR,
				Message: "Failed to save submission",
			}, err
		}
	}

	// Step 7: Return the submission response
	return &pb.SubmissionResponse{
		Status:  pb.SubmissionResponse_OK,
		Message: message,
		Passed:  int32(passedCount),
		Failed:  int32(failedCount),
	}, nil
}
