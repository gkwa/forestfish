package file

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
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

func DirExists(path string) bool {
	fileInfo, err := os.Stat(path)

	if err == nil {
		if fileInfo.IsDir() {
			return true
		} else {
			slog.Error("not a directory", "path", path)
			return false
		}
	} else if os.IsNotExist(err) {
		slog.Debug("directory does not exist", "path", path)
		return false
	} else {
		slog.Error("error checking directory", "path", path, err.Error())
		return false
	}
}

func IsPortOpen(host string, port int) bool {
	args := []string{"-z", host, strconv.Itoa(port), "-G", "5"}
	cmd := exec.Command("/usr/bin/nc", args...)
	_, err := cmd.Output()
	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			return false
		}
	}
	return true
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

// obsolete, use RunCmd instead
func CmdRun(cmd *exec.Cmd, cwd, stdOutLog, stdErrLog string) {
	// obsolete, use RunCmd instead
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

	slog.Debug("command output", "cmd", cmd.String(), "output", outStr)

	slog.Error("command error", "cmd", cmd.String(), "error", errStr)

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
