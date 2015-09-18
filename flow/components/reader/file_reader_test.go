package reader

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestNewInstanceHavingInputFile(t *testing.T) {
	// given
	file := new(os.File)
	c := newFileReaderMockedConfiguration(file)

	// when
	sut := newFileReader(c)

	// then
	assert.Equal(t, file, sut.inputFile, "Expected file is different than the specified in configuration.")
}

func TestPipe(t *testing.T) {

}
