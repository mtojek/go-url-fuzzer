package reader

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/stretchr/testify/assert"
)

func TestNewInstanceHasInputFile(t *testing.T) {
	assert := assert.New(t)

	// given
	file := new(os.File)
	builder := configuration.NewBuilder()
	configuration := builder.FuzzSetFile(file).Build()

	// when
	sut := newFileReader(configuration)

	// then
	assert.Equal(file, sut.inputFile, "Expected file is different than the specified in configuration.")
}

func TestPipe(t *testing.T) {
	assert := assert.New(t)

	// given
	file, error := openFuzzSetFile("fuzz_01.txt")
	assert.Nil(error, "Fuzz set file must be available, error: %v", error)

	builder := configuration.NewBuilder()
	configuration := builder.FuzzSetFile(file).Build()
	sut := newFileReader(configuration)

	var out = make(chan string, 3) // number of lines in fuzz file
	var done = make(chan bool, 1)

	// when
	sut.pipe(out, done)

	// then
	assert.True(<-done, "Reading file should be finished.")
	assert.Equal("001", <-out, "Invalid line read from file.")
	assert.Equal("002", <-out, "Invalid line read from file.")
	assert.Equal("003", <-out, "Invalid line read from file.")
	assert.Len(out, 0, "Out buffer should be empty now.")
	close(out)

	assert.Equal(^(uintptr(0)), file.Fd(), "File descriptor should be closed now.")
}

func TestCloseFile(t *testing.T) {
	assert := assert.New(t)

	// given
	file, error := openFuzzSetFile("fuzz_01.txt")
	assert.Nil(error, "Fuzz set file must be available, error: %v", error)

	builder := configuration.NewBuilder()
	configuration := builder.FuzzSetFile(file).Build()
	sut := newFileReader(configuration)

	// when
	sut.closeFile()

	// then
	assert.Equal(^(uintptr(0)), file.Fd(), "File descriptor should be closed now.")
}

func openFuzzSetFile(fileName string) (*os.File, error) {
	fuzzSetFile := filepath.Join("../../../resources/input-data/", fileName)
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
