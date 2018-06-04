package cmd

import (
	"bytes"
	"testing"

	shellwords "github.com/mattn/go-shellwords"
)

func TestShow(t *testing.T) {
	// table driven test
	cases := []struct {
		command string
		want    string
		errWant string
	}{
		{command: "cmd-test show", want: "", errWant: "Parameter error: Optstr is required"},
		{command: "cmd-test show --int 10", want: "", errWant: "Parameter error: Optstr is required"},
		{command: "cmd-test show --str test", want: "show called: optint: 0, optstr: test \n"},
		{command: "cmd-test show --str \"test1 test2\"", want: "", errWant: "Parameter error: Optstr is not valid test1 test2"},
		{command: "cmd-test show --int 1000 --str 123", want: "", errWant: "Parameter error: Optint cannot be greater than 10"},
		{command: "cmd-test show --int -1 --str abc", want: "", errWant: "Parameter error: Optint must be greater than 0"},
		{command: "cmd-test show --int 1 --str abc", want: "show called: optint: 1, optstr: abc \n"},
	}

	for _, c := range cases {
		t.Run(c.command, func(t *testing.T) {
			buf := new(bytes.Buffer)
			rootCmd := NewCmdRoot()
			rootCmd.SetOutput(buf)
			cmdArgs, err := shellwords.Parse(c.command)
			// cmdArgs := strings.Split(c.command, " ")
			if err != nil {
				t.Fatalf("args parse error: %+v\n", err)
			}
			rootCmd.SetArgs(cmdArgs[1:])
			if err := rootCmd.Execute(); err != nil {
				if c.errWant != err.Error() {
					t.Errorf("unexpected error response: errWant:`%+v`, get:`%+v`", c.errWant, err.Error())
				}
			}

			get := buf.String()
			if c.want != get {
				t.Errorf("unexpected response: want:`%+v`, get:`%+v`", c.want, get)
			}
		})
	}
}
