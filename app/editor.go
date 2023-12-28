// Package app edit in environment editor func
package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// getEnvEditor returns EDITOR or VISUAL environment variable or an error if not defined.
func getEnvEditor() (string, error) {
	editor := ""

	for _, e := range []string{"EDITOR", "VISUAL"} {
		editor = os.Getenv(e)
		if editor != "" {
			break
		}
	}

	if editor == "" {
		return "", fmt.Errorf("no value for environment variables were found")
	}

	return editor, nil
}

func editInEnvEditor(title, content string) (string, error) {

	// Try EDITOR then VISUAL

	editor, err := getEnvEditor()
	if err != nil {
		return "", err
	}

	// Create temp file name
	tmpFile, err := ioutil.TempFile("", title)
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	// Put text in file
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		return "", err
	}

	// Look up editor path
	cmd, err := exec.LookPath(editor)
	if err != nil {
		return "", err
	}

	// Define process attributes
	var attr os.ProcAttr
	attr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}

	// Start the editor
	p, err := os.StartProcess(cmd, []string{cmd, tmpFile.Name()}, &attr)
	if err != nil {
		return "", err
	}
	_, _ = p.Wait()

	// Read file content
	text, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(text)), nil
}
