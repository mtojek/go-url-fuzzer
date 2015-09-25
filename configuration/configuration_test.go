package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadersAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	headerName := "a_header_name"
	headerValue := "a_header_value"
	expectedHeaders := map[string]string{headerName: headerValue}

	builder := NewBuilder()
	sut := builder.Headers(expectedHeaders).Build()

	// when
	actualHeaders, exists := sut.Headers()

	// then
	assert.Equal(actualHeaders, expectedHeaders, "Assgined different headers.")
	assert.True(exists, "Value should exist")
}

func TestHeadersNotAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := newConfiguration()

	// when
	actualHeaders, exists := sut.Headers()

	// then
	assert.Empty(actualHeaders, "Headers should be empty.")
	assert.False(exists, "Value should not exist")
}

func TestOutputFileAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	expectedOutputFile := "output_file.ext"

	builder := NewBuilder()
	sut := builder.OutputFile(expectedOutputFile).Build()

	// when
	actualOutputFile, defined := sut.OutputFile()

	// then
	assert.Equal(actualOutputFile, expectedOutputFile, "Assgined different output files.")
	assert.True(defined, "Value should be defined")
}

func TestOutputFileNotAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := newConfiguration()

	// when
	actualOutputFile, defined := sut.OutputFile()

	// then
	assert.Empty(actualOutputFile, "Output file should be empty.")
	assert.False(defined, "Value should not be defined")
}
