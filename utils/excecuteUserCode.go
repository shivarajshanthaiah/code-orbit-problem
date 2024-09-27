package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Write user code to a temporary file and run it
func ExecuteUserCode(code, input, tempFilePath string) (string, error) {
	// Write the user's code to a temporary Go file
	if err := ioutil.WriteFile(tempFilePath, []byte(code), 0644); err != nil {
		return "", fmt.Errorf("failed to write code to file: %v", err)
	}

	defer os.Remove(tempFilePath) // Ensure file is removed after execution

	// Run the code using the runCodeFromFile function
	return runCodeFromFile(tempFilePath, input)
}

// Generic function to handle the execution of Go code from a file
func runCodeFromFile(codeFile, input string) (string, error) {
	resultChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				errChan <- fmt.Errorf("panic occurred: %v", r)
			}
		}()

		cmd := exec.Command("go", "run", codeFile)
		cmd.Stdin = strings.NewReader(input)
		output, err := cmd.CombinedOutput()
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- string(output)
	}()

	select {
	case output := <-resultChan:
		return output, nil
	case err := <-errChan:
		return "", err
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("execution timed out")
	}
}