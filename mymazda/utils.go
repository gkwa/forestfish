package file

import (
	"log/slog"
	"os"
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
