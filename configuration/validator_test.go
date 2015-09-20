package configuration

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/mtojek/go-url-fuzzer/testutils/localserver"
	"github.com/stretchr/testify/assert"
)

func TestInvalidHeaders(t *testing.T) {
	assert := assert.New(t)

	// given
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": ""}
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
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "PUT", "OPTIONS"}
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
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var zero uint64
	configuration.workersNumber = &zero
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
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var thousand uint64 = 1000
	configuration.workersNumber = &thousand
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validateOffline()

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", tooManyWorkersError), "tooManyWorkersError should be returned.")
}

func TestRelativeBaseUrl(t *testing.T) {
	assert := assert.New(t)

	// given
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var one uint64 = 1
	configuration.workersNumber = &one
	relativeURL, _ := url.Parse("relative/url/1/2/3")
	configuration.baseURL = &relativeURL
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
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var one uint64 = 1
	configuration.workersNumber = &one
	relativeURL, _ := url.Parse("http://relative/url/1/2/3")
	configuration.baseURL = &relativeURL
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
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var one uint64 = 1
	configuration.workersNumber = &one
	invalidURL, _ := url.Parse("http://invalid-domain.tld/")
	configuration.baseURL = &invalidURL
	var responseTimeout = 1 * time.Second
	configuration.urlResponseTimeout = &responseTimeout
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
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var one uint64 = 1
	configuration.workersNumber = &one
	invalidURL, _ := url.Parse("invalid-domain.tld/test-dir/")
	configuration.baseURL = &invalidURL
	var responseTimeout = 1 * time.Second
	configuration.urlResponseTimeout = &responseTimeout
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

	configuration := newConfiguration()
	invalidURL, _ := url.Parse(scheme + "://" + hostPort)
	configuration.baseURL = &invalidURL
	var responseTimeout = 1 * time.Second
	configuration.urlResponseTimeout = &responseTimeout
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

	configuration := newConfiguration()
	invalidURL, _ := url.Parse(scheme + "://" + hostPort)
	configuration.baseURL = &invalidURL
	var responseTimeout = 1 * time.Second
	configuration.urlResponseTimeout = &responseTimeout
	sut := newValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	server.StartTLS("../resources/certs/server_ca.pem", "../resources/certs/server_ca.key")

	// when
	error := sut.validateOnline()

	server.Stop()

	// then
	assert.Nil(error, "There should not be error returned.")
}
