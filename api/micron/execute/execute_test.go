package execute

import (
	"testing"
)

func TestWithOutput(t *testing.T) {
	var output string
	var err error

	output, err = WithOutput("pwsh.exe", []string{"-c", "echo hello"})
	if err != nil {
		t.Errorf("execute valid with output: %v", err)
	}
	if output != "hello\r\n" {
		t.Errorf("execute valid with output not as expected")
	}

	output, err = WithOutput("NotExists.exe", []string{"-c", "echo hello"})
	if err == nil {
		t.Errorf("execute invalid with output: %v", err)
	}
	if output != "" {
		t.Errorf("execute invalid with output not as expected")
	}
}

func TestWithDetach(t *testing.T) {
	var err error

	err = WithDetach("pwsh.exe", []string{"-c", "echo hello"})
	if err != nil {
		t.Errorf("execute valid detached: %v", err)
	}

	err = WithDetach("NotExists.exe", []string{"-c", "echo hello"})
	if err == nil {
		t.Errorf("execute invalid detached should failed")
	}
}
