package problem

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/utils"
)

// RunMathProblem - Function to execute a mathematical problem
func RunMathProblem(code, input string) (string, error) {
	log.Println("this is math")
	tempFilePath := "./pkg/userCodeExcecution/temp/submission.go"
	return utils.ExecuteUserCode(code, input, tempFilePath)
}

// RunStringProblem - Function to execute a string manipulation problem
func RunStringProblem(code, input string) (string, error) {
	log.Println("this is string")
	tempFilePath := "./pkg/userCodeExcecution/temp/submission.go"
	return utils.ExecuteUserCode(code, input, tempFilePath)
}

func RunArrayProblem(code, input string) (string, error) {
	log.Println("this is array")
	tempFilePath := "./pkg/userCodeExcecution/temp/submission.go"
	return utils.ExecuteUserCode(code, input, tempFilePath)
}
