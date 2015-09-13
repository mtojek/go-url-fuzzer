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

	sut := newConfiguration()
	sut.headers = &expectedHeaders

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

	sut := newConfiguration()
	sut.outputFile = &expectedOutputFile

	// when
	actualOutputFile, exists := sut.OutputFile()

	// then
	assert.Equal(actualOutputFile, expectedOutputFile, "Assgined different output files.")
	assert.True(exists, "Value should exist")
}

func TestOutputFileNotAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := newConfiguration()

	// when
	actualOutputFile, exists := sut.OutputFile()

	// then
	assert.Empty(actualOutputFile, "Output file should be empty.")
	assert.False(exists, "Value should not exist")
}

func TestReportDirectoryAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	expectedReportDirectory := "report_directory"

	sut := newConfiguration()
	sut.reportDirectory = &expectedReportDirectory

	// when
	actualReportDirectory, exists := sut.ReportDirectory()

	// then
	assert.Equal(actualReportDirectory, expectedReportDirectory, "Assgined different report directories.")
	assert.True(exists, "Value should exist")
}

func TestReportDirectoryNotAvailable(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := newConfiguration()

	// when
	actualReportDirectory, exists := sut.ReportDirectory()

	// then
	assert.Empty(actualReportDirectory, "Report directory should be empty.")
	assert.False(exists, "Value should not exist")
}
