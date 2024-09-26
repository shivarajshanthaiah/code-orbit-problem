package usercodeexcecution

import (
	"fmt"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/userCodeExcecution/problem"
)

// RunUserCode - Entry point to execute the user's code based on the problem type
func RunUserCode(problemType, code, input string) (string, error) {
	switch problemType {
	case "string":
		return problem.RunStringProblem(code, input)
	case "math":
		return problem.RunMathProblem(code, input)
	default:
		return "", fmt.Errorf("unsupported problem type: %s", problemType)
	}
}
