package filez

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ibrt/golang-bites/internal"
)

// MustAbs is like filepath.Abs, but panics on error.
func MustAbs(path string) string {
	path, err := filepath.Abs(path)
	internal.MaybePanic(err)
	return path
}

// MustGetWD is like os.Getwd, but panics on error.
func MustGetWD() string {
	wd, err := os.Getwd()
	internal.MaybePanic(err)
	return wd
}

// MustUserHomeDir is like os.UserHomeDir, but panics on error.
func MustUserHomeDir() string {
	dirPath, err := os.UserHomeDir()
	internal.MaybePanic(err)
	return dirPath
}

// MustReadFile reads a file, panics on error.
func MustReadFile(filePath string) []byte {
	buf, err := ioutil.ReadFile(filePath)
	internal.MaybePanic(err)
	return buf
}

// MustWriteFile creates a file with the given mode and contents, also ensuring the containing folder exists.
func MustWriteFile(filePath string, dirMode os.FileMode, fileMode os.FileMode, contents []byte) string {
	internal.MaybePanic(os.MkdirAll(filepath.Dir(filePath), dirMode))
	internal.MaybePanic(ioutil.WriteFile(filePath, contents, fileMode))
	return filePath
}

// MustWriteTempFile creates a temporary file with the given contents.
func MustWriteTempFile(pattern string, contents []byte) string {
	fd, err := ioutil.TempFile("", pattern)
	internal.MaybePanic(err)
	defer func() {
		internal.MaybePanic(fd.Close())
	}()

	_, err = io.Copy(fd, bytes.NewReader(contents))
	internal.MaybePanic(err)
	return fd.Name()
}

// WithMustWriteTempFile calls f passing it the path to a new temporary file, which is wiped after f returns.
func WithMustWriteTempFile(pattern string, contents []byte, f func(filePath string)) {
	filePath := MustWriteTempFile(pattern, contents)
	defer func() {
		internal.MaybePanic(os.RemoveAll(filePath))
	}()

	f(filePath)
}

// MustPrepareDir deletes the given directory and its contents (if present) and recreates it.
func MustPrepareDir(dirPath string, dirMode os.FileMode) {
	internal.MaybePanic(os.RemoveAll(dirPath))
	internal.MaybePanic(os.MkdirAll(dirPath, dirMode))
}

// MustCreateTempDir is like os.MkdirTemp, but panics on error.
func MustCreateTempDir(pattern string) string {
	dirPath, err := os.MkdirTemp("", pattern)
	internal.MaybePanic(err)
	return dirPath
}

// WithMustCreateTempDir calls f passing it the path to a new temporary directory, which is wiped after f returns.
func WithMustCreateTempDir(pattern string, f func(dirPath string)) {
	dirPath := MustCreateTempDir(pattern)
	defer func() {
		internal.MaybePanic(os.RemoveAll(dirPath))
	}()

	f(dirPath)
}

// MustCheckExists checks if the given path exists, panics on error (other than os.ErrNotExist).
func MustCheckExists(fileOrDirPath string) bool {
	_, err := os.Stat(fileOrDirPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
