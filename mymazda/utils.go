package file

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
	}
}

func CreateFile(p string) *os.File {
	// https://gobyexample.com/defer
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func CloseFile(f *os.File) {
	// https://gobyexample.com/defer
	err := f.Close()
	if err != nil {
		slog.Error("file close", "error", err.Error())
		os.Exit(1)
	}
}

func CmdRun(cmd *exec.Cmd, cwd, stdOutLog, stdErrLog string) {
	cmd.Dir = cwd

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	slog.Debug("running command", "cmd", cmd.String(), "cwd", cwd)
	err := cmd.Run()
	if err != nil {
		slog.Error("error running command", "cmd", cmd.String(), "error", err.Error())
	}

	outStr, errStr := stdout.String(), stderr.String()

	if stdout.Len() > 0 {
		f := CreateFile(stdOutLog)
		defer CloseFile(f)

		f.WriteString(outStr)
	}

	if stderr.Len() > 0 {
		f := CreateFile(stdErrLog)
		defer CloseFile(f)

		f.WriteString(errStr)
	}
}
