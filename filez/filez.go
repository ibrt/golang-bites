package filez

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// MustAbs is like filepath.Abs, but panics on error.
func MustAbs(path string) string {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return path
}

// MustGetWD is like os.Getwd, but panics on error.
func MustGetWD() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

// MustUserHomeDir is like os.UserHomeDir, but panics on error.
func MustUserHomeDir() string {
	dirPath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return dirPath
}

// MustReadFile reads a file, panics on error.
func MustReadFile(filePath string) []byte {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return buf
}

// MustWriteFile creates a file with the given mode and contents, also ensuring the containing folder exists.
func MustWriteFile(filePath string, dirMode os.FileMode, fileMode os.FileMode, contents []byte) string {
	if err := os.MkdirAll(filepath.Dir(filePath), dirMode); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(filePath, contents, fileMode); err != nil {
		panic(err)
	}

	return filePath
}

// MustWriteTempFile creates a temporary file with the given contents.
func MustWriteTempFile(pattern string, contents []byte) string {
	fd, err := ioutil.TempFile("", pattern)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fd.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := io.Copy(fd, bytes.NewReader(contents)); err != nil {
		panic(err)
	}

	return fd.Name()
}

// MustPrepareDir deletes the given directory and its contents (if present) and recreates it.
func MustPrepareDir(dirPath string, dirMode os.FileMode) {
	if err := os.RemoveAll(dirPath); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(dirPath, dirMode); err != nil {
		panic(err)
	}
}

// MustCreateTempDir is like os.MkdirTemp, but panics on error.
func MustCreateTempDir(pattern string) string {
	dirPath, err := os.MkdirTemp("", pattern)
	if err != nil {
		panic(err)
	}
	return dirPath
}

// WithMustCreateTempDir calls f passing it the path to a new temporary directory, which is wiped after f returns.
func WithMustCreateTempDir(pattern string, f func(dirPath string)) {
	dirPath := MustCreateTempDir(pattern)
	defer func() {
		if err := os.RemoveAll(dirPath); err != nil {
			panic(err)
		}
	}()

	f(dirPath)
}
