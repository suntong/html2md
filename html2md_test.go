package main_test

import (
	"os"
	"strings"
	"testing"
)

const (
	cmdTest = "./html2md"
	dirTest = "."

	boldText   = "<strong>Bold Text</strong>"
	boldEscape = "<strong>option src_ip</strong>"
)

type testCase struct {
	name, out, in string
	args          []string
}

func TestExec(t *testing.T) {
	testData := []testCase{
		{
			"BoldText", "**Bold Text**", boldText, []string{"-i"},
		},
		{
			"BoldText-delimiter", "__Bold Text__", boldText,
			[]string{"-i", "--opt-strong-delimiter", "__"},
		},
		{
			"BoldEscape", "**option src\\_ip**", boldEscape, []string{"-i"},
		},
		{
			"Checkbox", "- [x] Checked!\n- [ ] Check Me!",
			"<ul><li><input type=checkbox checked>Checked!</li><li><input type=checkbox>Check Me!</li></ul>",
			[]string{"-i", "-G"},
		},
		{
			"PluginStrikethrough", "Only ~~blue ones~~ ~~left~~",
			"Only <del>blue ones</del> <s> left</s>",
			[]string{"-i", "--plugin-strikethrough"},
		},
		// {
		// 	"", "", "", []string{"-i"},
		// },
	}

	os.Chdir(dirTest)

	t.Logf("\n\n== Testing options and plugins\n\n")

	for _, tc := range testData {
		r := strings.TrimSpace(testIt(t, tc.name, tc.in, tc.args...))
		if tc.out != r {
			t.Errorf(`expected "%s" but got "%s" instead`, tc.out, r)
		}
	}
}
