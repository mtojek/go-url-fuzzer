package configuration

import (
	"testing"

	"net/url"
	"os"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddHeaders(t *testing.T) {
	assert := assert.New(t)

	// given
	headers := map[string]string{"Set-Cookie": "b=1"}
	sut := NewBuilder()

	// when
	result := sut.Headers(headers).Build()

	// then
	assert.Equal(result.headers, &headers, "HTTP headers should be equal")
}

func TestAddMethods(t *testing.T) {
	assert := assert.New(t)

	// given
	methods := []string{"GET", "POST"}
	sut := NewBuilder()

	// when
	result := sut.Methods(methods).Build()

	// then
	assert.Equal(result.methods, &methods, "HTTP methods should be equal")
}

func TestAddOutputFile(t *testing.T) {
	assert := assert.New(t)

	// given
	outputFile := "output-file"
	sut := NewBuilder()

	// when
	result := sut.OutputFile(outputFile).Build()

	// then
	assert.Equal(result.outputFile, &outputFile, "Output files should be equal")
}

func TestAddURLResponseTimeout(t *testing.T) {
	assert := assert.New(t)

	// given
	urlResponseTimeout := 8 * time.Microsecond
	sut := NewBuilder()

	// when
	result := sut.URLResponseTimeout(urlResponseTimeout).Build()

	// then
	assert.Equal(result.urlResponseTimeout, &urlResponseTimeout, "URL response timeouts should be equal")
}

func TestAddHTTPErrorCode(t *testing.T) {
	assert := assert.New(t)

	// given
	var httpErrorCode uint64 = 492
	sut := NewBuilder()

	// when
	result := sut.HTTPErrorCode(httpErrorCode).Build()

	// then
	assert.Equal(result.httpErrorCode, &httpErrorCode, "HTTP error codes should be equal")
}

func TestAddWorkersNumber(t *testing.T) {
	assert := assert.New(t)

	// given
	var workersNumber uint64 = 15
	sut := NewBuilder()

	// when
	result := sut.WorkersNumber(workersNumber).Build()

	// then
	assert.Equal(result.workersNumber, &workersNumber, "Numbers of workers should be equal")
}

func TestAddWorkerWaitPeriod(t *testing.T) {
	assert := assert.New(t)

	// given
	workerWaitPeriod := 9 * time.Hour
	sut := NewBuilder()

	// when
	result := sut.WorkerWaitPeriod(workerWaitPeriod).Build()

	// then
	assert.Equal(result.workerWaitPeriod, &workerWaitPeriod, "Worker wait periods should be equal")
}

func TestAddFuzzSetFile(t *testing.T) {
	assert := assert.New(t)

	// given
	fuzzSetFile := new(os.File)
	sut := NewBuilder()

	// when
	result := sut.FuzzSetFile(fuzzSetFile).Build()

	// then
	assert.Equal(result.fuzzSetFile, &fuzzSetFile, "Fuzz set files should be equal")
}

func TestAddBaseURL(t *testing.T) {
	assert := assert.New(t)

	// given
	baseURL := new(url.URL)
	sut := NewBuilder()

	// when
	result := sut.BaseURL(baseURL).Build()

	// then
	assert.Equal(result.baseURL, &baseURL, "Base URLs should be equal")
}

func TestBuild(t *testing.T) {
	assert := assert.New(t)

	// given
	sut := NewBuilder()

	// when
	result := sut.Build()

	// then
	assert.NotNil(result, "Built configuration instance should not be nil")
}
