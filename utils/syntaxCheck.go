package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// func CheckSyntax(code string) error {
// 	// Create a temporary file
// 	tmpfile, err := ioutil.TempFile("", "example*.go")
// 	if err != nil {
// 		return err
// 	}
// 	defer os.Remove(tmpfile.Name())

// 	// Write the code to the temporary file
// 	if _, err := tmpfile.Write([]byte(code)); err != nil {
// 		return err
// 	}
// 	if err := tmpfile.Close(); err != nil {
// 		return err
// 	}

// 	// Try to compile the code
// 	cmd := exec.Command("go", "build", tmpfile.Name())
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return fmt.Errorf("compilation error: %v\n%s", err, string(output))
// 	}

// 	return nil
// }

func CheckSyntax(code string) error {
	// Define the path to the existing temp directory
	tmpDir := "./pkg/userCodeExcecution/temp"

	// Create a unique temporary file within the temp directory
	tmpfilePath := filepath.Join(tmpDir, "syntax_check.go")
	if err := ioutil.WriteFile(tmpfilePath, []byte(code), 0644); err != nil {
		return fmt.Errorf("failed to write code to temporary file: %v", err)
	}

	// Ensure the temporary file is removed after execution
	defer os.Remove(tmpfilePath)

	// Try to compile the code
	cmd := exec.Command("go", "build", "-o", filepath.Join(tmpDir, "main"), tmpfilePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("compilation error: %v\n%s", err, string(output))
	}

	// Optionally remove the compiled binary after checking syntax
	defer os.Remove(filepath.Join(tmpDir, "main"))

	return nil
}
