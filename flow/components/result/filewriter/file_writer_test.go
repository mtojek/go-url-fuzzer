package filewriter

import (
	"testing"

	"io/ioutil"
	"log"
	"os"

	"github.com/mtojek/go-url-fuzzer/configuration"
	"github.com/mtojek/go-url-fuzzer/flow/messages"
	"github.com/stretchr/testify/assert"
)

func TestNewFileWriterOutputFileUndefined(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := configuration.NewBuilder()
	configuration := builder.Build()

	// when
	sut := NewFileWriter(configuration)

	// then
	assert.NotNil(sut, "File writer is not nil")
	assert.Nil(sut.outputFile, "Output file has not been given")
}

func TestNewFileWriterOutputFileDefined(t *testing.T) {
	assert := assert.New(t)

	// given
	aFile, error := ioutil.TempFile("", "TestNewFileWriterOutputFileDefined")
	if nil != error {
		log.Fatalf("Can't create a temporary output file: %v", error)
	}
	path := aFile.Name()

	builder := configuration.NewBuilder()
	configuration := builder.
		OutputFile(path).
		Build()

	// when
	sut := NewFileWriter(configuration)
	defer cleanAndClose(aFile)

	fi, error := sut.outputFile.Stat()
	if nil != error {
		log.Fatalf("Error occured while checking size of file \"%v\": %v", path, error)
	}

	// then
	assert.NotNil(sut.outputFile, "Output file has been defined")
	assert.True(fi.Size() == 0, "File should be empty")
}

func TestOnFoundEntry(t *testing.T) {
	assert := assert.New(t)

	// given
	aFile, error := ioutil.TempFile("", "TestNewFileWriterOutputFileDefined")
	if nil != error {
		log.Fatalf("Can't create a temporary output file: %v", error)
	}
	path := aFile.Name()

	builder := configuration.NewBuilder()
	configuration := builder.
		OutputFile(path).
		Build()
	firstFoundEntry := messages.NewFoundEntry("1", "2", 3)
	sut := NewFileWriter(configuration)

	// when
	sut.OnFoundEntry(firstFoundEntry)
	defer cleanAndClose(aFile)

	fi, error := sut.outputFile.Stat()
	if nil != error {
		log.Fatalf("Error occured while checking size of file \"%v\": %v", path, error)
	}

	// then
	assert.True(fi.Size() > 0, "File should have contents")
}

func TestFinish(t *testing.T) {
	assert := assert.New(t)

	// given
	aFile, error := ioutil.TempFile("", "TestNewFileWriterOutputFileDefined")
	if nil != error {
		log.Fatalf("Can't create a temporary output file: %v", error)
	}
	path := aFile.Name()

	builder := configuration.NewBuilder()
	configuration := builder.
		OutputFile(path).
		Build()
	sut := NewFileWriter(configuration)

	// when
	fi, error := sut.outputFile.Stat()
	if nil != error {
		log.Fatalf("Error occured while checking size of file \"%v\": %v", path, error)
	}

	sut.Finish()
	defer clean(path)

	// then
	assert.NotNil(sut.outputFile, "Output file has been defined")
	assert.True(fi.Size() == 0, "File should be empty")
	assert.Equal(^(uintptr(0)), sut.outputFile.Fd(), "File descriptor should be closed now.")
}

func cleanAndClose(aFile *os.File) {
	path := aFile.Name()
	if error := aFile.Close(); nil != error {
		log.Fatalf("Error occured while closing file \"%v\": %v", path, error)
	}

	clean(path)
}

func clean(path string) {
	if error := os.Remove(path); nil != error {
		log.Fatalf("Error occured while removing file \"%v\": %v", path, error)
	}
}
