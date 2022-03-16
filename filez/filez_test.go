package filez_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ibrt/golang-bites/filez"
	"github.com/ibrt/golang-bites/internal"
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
	var tempFilePath string

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

	filez.WithMustWriteTempFile("golang-bites", []byte("temp"), func(filePath string) {
		tempFilePath = filePath
		require.Equal(t, []byte("temp"), filez.MustReadFile(filePath))
	})

	_, err = os.Stat(tempFilePath)
	require.Error(t, err)
	require.True(t, os.IsNotExist(err))
}

func TestMustCheckExists(t *testing.T) {
	filez.WithMustCreateTempDir("golang-bites", func(dirPath string) {
		require.True(t, filez.MustCheckExists(dirPath))
	})

	var deletedFilePath string
	filez.WithMustWriteTempFile("golang-bites", []byte("temp"), func(filePath string) {
		require.True(t, filez.MustCheckExists(filePath))
		deletedFilePath = filePath
	})

	require.False(t, filez.MustCheckExists(deletedFilePath))

	require.Panics(t, func() {
		filez.MustCheckExists(string([]byte{0}))
	})
}

func TestMustCopyEmbedFSSimple(t *testing.T) {
	filez.WithMustCreateTempDir("golang-bites", func(dirPath string) {
		filez.MustCopyEmbedFSSimple(internal.ExampleDirAssetFS, internal.ExampleDirPathPrefix, dirPath)

		paths := make([]string, 0)
		require.NoError(t, filepath.WalkDir(dirPath, func(path string, _ fs.DirEntry, err error) error {
			require.NoError(t, err)
			paths = append(paths, strings.TrimPrefix(path, dirPath))
			return nil
		}))

		require.Equal(t, []string{"", "/child-dir", "/child-dir/second.txt", "/first.txt"}, paths)
		require.Equal(t, []byte("FIRST"), filez.MustReadFile(filepath.Join(dirPath, "first.txt")))
		require.Equal(t, []byte("SECOND"), filez.MustReadFile(filepath.Join(dirPath, "child-dir", "second.txt")))
	})
}
