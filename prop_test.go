package main_test

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

// testIt runs @cmdTest with @argv and @in as stdin in @dirTest
// then return the output
func testIt(t *testing.T, name, in string, argv ...string) string {
	var (
		appOut bytes.Buffer
		cmd    = exec.Command(cmdTest, argv...)
	)

	t.Logf("Testing '%s':\n\t%s %s\n", name, cmdTest, strings.Join(argv, " "))

	cmd.Stdin = strings.NewReader(in)
	cmd.Stdout = &appOut

	err := cmd.Start()
	if err != nil {
		t.Errorf("start error [%s: %s] %s.", name, argv, err)
	}
	err = cmd.Wait()
	if err != nil {
		t.Errorf("exit error [%s: %s] %s.", name, argv, err)
	}

	return appOut.String()
}
