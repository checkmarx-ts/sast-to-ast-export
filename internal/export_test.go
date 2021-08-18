package internal

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	prefix = "ast-sast-export-test-export"
)

func TestCreateExport(t *testing.T) {
	export, err := CreateExport(prefix)
	assert.NoError(t, err)
	defer export.Clean()

	info, statErr := os.Stat(export.TmpDir)
	assert.NoError(t, statErr)
	assert.True(t, info.IsDir())
	assert.Contains(t, export.TmpDir, prefix)
}

func TestExport_AddFile(t *testing.T) {
	export, err := CreateExport(prefix)
	assert.NoError(t, err)
	defer export.Clean()

	addErr := export.AddFile("test1.txt", []byte("this is test1"))
	assert.NoError(t, addErr)

	expectedFileList := []string{"test1.txt"}
	assert.Equal(t, expectedFileList, export.FileList)

	test1FileName := path.Join(export.TmpDir, "test1.txt")
	info, statErr := os.Stat(test1FileName)
	assert.NoError(t, statErr)
	assert.False(t, info.IsDir())

	content, ioErr := ioutil.ReadFile(test1FileName)
	assert.NoError(t, ioErr)
	assert.Equal(t, "this is test1", string(content))
}

func TestExport_CreateExportPackage(t *testing.T) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), prefix)
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	export, err := CreateExport(prefix)
	assert.NoError(t, err)
	defer export.Clean()

	addErr1 := export.AddFile("test1.txt", []byte("this is test1"))
	assert.NoError(t, addErr1)

	addErr2 := export.AddFile("test2.txt", []byte("this is test2"))
	assert.NoError(t, addErr2)

	exportFileName, exportErr := export.CreateExportPackage(prefix, tmpDir)
	assert.NoError(t, exportErr)

	info, statErr := os.Stat(exportFileName)
	assert.NoError(t, statErr)
	assert.False(t, info.IsDir())
	assert.Contains(t, exportFileName, prefix)

	chdirErr := os.Chdir(tmpDir)
	assert.NoError(t, chdirErr)

	zipReader, zipErr := zip.OpenReader(exportFileName)
	assert.NoError(t, zipErr)

	encryptedKeyFile, zipErr := zipReader.Open(EncryptedKeyFileName)
	assert.NoError(t, zipErr)

	_, keyStatErr := encryptedKeyFile.Stat()
	assert.NoError(t, keyStatErr)

	encryptedZipFile, zipErr := zipReader.Open(EncryptedZipFileName)
	assert.NoError(t, zipErr)

	_, zipStatErr := encryptedZipFile.Stat()
	assert.NoError(t, zipStatErr)
}

func TestExport_Clean(t *testing.T) {
	export, err := CreateExport(prefix)
	assert.NoError(t, err)

	cleanErr := export.Clean()
	assert.NoError(t, cleanErr)

	_, statErr := os.Stat(export.TmpDir)
	assert.Error(t, statErr)
	assert.True(t, os.IsNotExist(statErr))
}

func TestCreateExportFileName(t *testing.T) {
	now := time.Date(2021, time.August, 18, 12, 27, 34, 0, time.UTC)

	result := CreateExportFileName(prefix, now)

	expected := fmt.Sprintf("%s-2021-08-18-12-27-34.zip", prefix)
	assert.Equal(t, expected, result)
}
