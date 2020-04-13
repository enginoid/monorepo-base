package gcloud

import (
	"bytes"
	"fmt"
	"os/exec"
)

// simpleStringExec executes a command and returns the entire output as a
// string. If it fails, it returns the status code and contents of stderr
// as the error message.
//
// This increases readability of code calling simple commands that return
// output but don't require multiple cases of error handling.
func simpleStringExec(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr)
	}
	return stdout.String(), nil
}
