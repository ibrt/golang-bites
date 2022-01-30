package filez_test

import (
	"github.com/ibrt/golang-bites/filez"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestMustAbs(t *testing.T) {
	require.NotPanics(t, func() {
		require.NotEqual(t, "path", filez.MustAbs("path"))
	})
}

func TestGetWD(t *testing.T) {
	require.NotPanics(t, func() {
		require.NotEmpty(t, filez.MustGetWD())
	})
}

func TestMustUserHomeDir(t *testing.T) {
	require.NotPanics(t, func() {
		require.NotEmpty(t, filez.MustUserHomeDir())
	})
}

func TestFilesAndDirectories(t *testing.T) {
	filePath := filez.MustWriteTempFile("golang-bites", []byte("temp"))
	defer func() {
		require.NoError(t, os.RemoveAll(filePath))
	}()
	require.Equal(t, []byte("temp"), filez.MustReadFile(filePath))

	var tempDirPath string

	filez.WithMustCreateTempDir("golang-bites", func(dirPath string) {
		tempDirPath = dirPath
		filePath := filepath.Join(dirPath, "parent", "child")
		filez.MustWriteFile(filePath, 0777, 0666, []byte("temp"))
		require.Equal(t, []byte("temp"), filez.MustReadFile(filePath))

		filez.MustPrepareDir(dirPath, 0777)
		_, err := os.Stat(filePath)
		require.Error(t, err)
		require.True(t, os.IsNotExist(err))
	})

	_, err := os.Stat(tempDirPath)
	require.Error(t, err)
	require.True(t, os.IsNotExist(err))
}
