package file

import (
	"log/slog"
	"os"
	"os/user"
	"strings"
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

func ExpandTilde(path string) (string, error) {
	if strings.HasPrefix(path, "~/") || path == "~" {
		currentUser, err := user.Current()
		if err != nil {
			return "", err
		}
		return strings.Replace(path, "~", currentUser.HomeDir, 1), nil
	}
	return path, nil
}
