package configuration

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/mtojek/localserver"
	"github.com/stretchr/testify/assert"
)

func TestInvalidHeaders(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := NewBuilder()
	configuration := builder.Headers(map[string]string{"a_header": ""}).Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", missingHeaderValueError), "missingHeaderValueError should be returned.")
}

func TestRepeatedMethods(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := NewBuilder()
	configuration := builder.
		Headers(map[string]string{"a_header": "a_value"}).
		Methods([]string{"PUT", "POST", "PUT", "OPTIONS"}).
		Build()

	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", repeatedHTTPMethodError), "repeatedHttpMethodError should be returned.")
}

func TestZeroWorkersNumber(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := NewBuilder()
	configuration := builder.
		Headers(map[string]string{"a_header": "a_value"}).
		Methods([]string{"POST", "PUT", "OPTIONS"}).
		WorkersNumber(0).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", zeroWorkersNumberError), "zeroWorkersNumberError should be returned.")
}

func TestTooManyWorkersNumber(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := NewBuilder()
	configuration := builder.
		Headers(map[string]string{"a_header": "a_value"}).
		Methods([]string{"POST", "PUT", "OPTIONS"}).
		WorkersNumber(1000).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", tooManyWorkersError), "tooManyWorkersError should be returned.")
}

func TestInvalidHTTPErrorCodeBeforeRange(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := NewBuilder()
	configuration := builder.
		Headers(map[string]string{"a_header": "a_value"}).
		Methods([]string{"POST", "PUT", "OPTIONS"}).
		WorkersNumber(1).
		HTTPErrorCode(99).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", invalidHTTPErrorCodeError), "invalidHTTPErrorCodeError should be returned.")
}

func TestInvalidHTTPErrorCodeAfterRange(t *testing.T) {
	assert := assert.New(t)

	// given
	builder := NewBuilder()
	configuration := builder.
		Headers(map[string]string{"a_header": "a_value"}).
		Methods([]string{"POST", "PUT", "OPTIONS"}).
		WorkersNumber(1).
		HTTPErrorCode(600).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", invalidHTTPErrorCodeError), "invalidHTTPErrorCodeError should be returned.")
}

func TestRelativeBaseUrl(t *testing.T) {
	assert := assert.New(t)

	// given
	relativeURL, _ := url.Parse("relative/url/1/2/3")

	builder := NewBuilder()
	configuration := builder.
		BaseURL(relativeURL).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", relativeBaseURLError), "relativeBaseUrlError should be returned.")
}

func TestValidOfflineConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	relativeURL, _ := url.Parse("http://relative/url/1/2/3")

	builder := NewBuilder()
	configuration := builder.
		Headers(map[string]string{"a_header": "a_value"}).
		Methods([]string{"POST", "PUT", "OPTIONS"}).
		WorkersNumber(1).
		HTTPErrorCode(500).
		BaseURL(relativeURL).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.Nil(error, "There should not be error returned.")
}

func TestInvalidHostBaseUrl(t *testing.T) {
	assert := assert.New(t)

	// given
	invalidURL, _ := url.Parse("http://invalid-domain.tld/")

	builder := NewBuilder()
	configuration := builder.
		BaseURL(invalidURL).
		URLResponseTimeout(1 * time.Second).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOnline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", unableToConnectToHostBaseURLError), "unableToConnectToHostBaseURLError should be returned.")
}

func TestMissingSchemeBaseUrl(t *testing.T) {
	assert := assert.New(t)

	// given
	invalidURL, _ := url.Parse("invalid-domain.tld/test-dir/")

	builder := NewBuilder()
	configuration := builder.
		BaseURL(invalidURL).
		URLResponseTimeout(1 * time.Second).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOnline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", unableToConnectToHostBaseURLError), "unableToConnectToHostBaseURLError should be returned.")
}

func TestValidHttpHostBaseUrl(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "127.0.0.1:10601"
	scheme := "http"
	server := localserver.NewLocalServer(hostPort, scheme)
	validURL, _ := url.Parse(scheme + "://" + hostPort)

	builder := NewBuilder()
	configuration := builder.
		BaseURL(validURL).
		URLResponseTimeout(1 * time.Second).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	server.Start()

	// when
	error := sut.validateOnline()
	server.Stop()

	// then
	assert.Nil(error, "There should not be error returned.")
}

func TestValidHttpsHostBaseUrl(t *testing.T) {
	assert := assert.New(t)

	// given
	hostPort := "localhost:10602"
	scheme := "https"
	server := localserver.NewLocalServer(hostPort, scheme)
	validURL, _ := url.Parse(scheme + "://" + hostPort)

	builder := NewBuilder()
	configuration := builder.
		BaseURL(validURL).
		URLResponseTimeout(1 * time.Second).
		Build()
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	server.StartTLS("../resources/certs/server_ca.pem", "../resources/certs/server_ca.key")

	// when
	error := sut.validateOnline()

	server.Stop()

	// then
	assert.Nil(error, "There should not be error returned.")
}
