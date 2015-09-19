package reader

import (
	"testing"

	"os"

	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func TestNewInstanceHavingInputFile(t *testing.T) {
	assert := assert.New(t)

	// given
	file := new(os.File)
	c := newFileReaderMockedConfiguration(file)

	// when
	sut := newFileReader(c)

	// then
	assert.Equal(file, sut.inputFile, "Expected file is different than the specified in configuration.")
}

func TestPipe(t *testing.T) {
	assert := assert.New(t)

	// given
	file, error := openFuzzSetFile("fuzz_01.txt")
	assert.Nil(error, "Fuzz set file must be available, error: %v", error)

	c := newFileReaderMockedConfiguration(file)
	sut := newFileReader(c)

	var in = make(chan string, 3) // number of lines in fuzz file
	var done = make(chan bool, 1)

	// when
	sut.pipe(in, done)

	// then
	assert.True(<-done, "Reading file should be finished.")
	assert.Equal("001", <-in, "Invalid line read from file.")
	assert.Equal("002", <-in, "Invalid line read from file.")
	assert.Equal("003", <-in, "Invalid line read from file.")
	assert.Len(in, 0, "Read buffer should empty now.")

	assert.Equal(^(uintptr(0)), file.Fd(), "File descriptor should be closed now.")
}

func TestCloseFile(t *testing.T) {
	assert := assert.New(t)

	// given
	file, error := openFuzzSetFile("fuzz_01.txt")
	assert.Nil(error, "Fuzz set file must be available, error: %v", error)

	c := newFileReaderMockedConfiguration(file)
	sut := newFileReader(c)

	// when
	sut.closeFile()

	// then
	assert.Equal(^(uintptr(0)), file.Fd(), "File descriptor should be closed now.")
}

func openFuzzSetFile(fileName string) (*os.File, error) {
	fuzzSetFile := filepath.Join("../../../input-data/", fileName)
	absoluteFilePath, _ := filepath.Abs(fuzzSetFile)

	file, error := os.OpenFile(absoluteFilePath, os.O_RDONLY, 0666)
	if nil != error {
		return nil, error
	}

	_, error = file.Stat()
	if nil != error {
		return nil, error
	}

	return file, nil
}
