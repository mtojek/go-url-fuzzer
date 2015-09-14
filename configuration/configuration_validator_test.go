package configuration

import (
	"testing"

	"fmt"

	"net/url"

	"github.com/stretchr/testify/assert"
)

func TestInvalidHeaders(t *testing.T) {
	assert := assert.New(t)

	// given
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": ""}
	sut := newConfigurationValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validate(nil)

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
	sut := newConfigurationValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validate(nil)

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
	sut := newConfigurationValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validate(nil)

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", zeroWorkersNumberError), "zeroWorkersNumberError should be returned.")
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
	sut := newConfigurationValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validate(nil)

	// then
	assert.NotNil(error, "There should be error returned.")
	assert.Equal(error.Error(), fmt.Sprintf("%v", relativeBaseURLError), "relativeBaseUrlError should be returned.")
}

func TestValidConfiguration(t *testing.T) {
	assert := assert.New(t)

	// given
	configuration := newConfiguration()
	configuration.headers = &map[string]string{"a_header": "a_value"}
	configuration.methods = &[]string{"PUT", "POST", "OPTIONS"}
	var one uint64 = 1
	configuration.workersNumber = &one
	relativeURL, _ := url.Parse("http://relative/url/1/2/3")
	configuration.baseURL = &relativeURL
	sut := newConfigurationValidator(configuration)
	sut.errorTagMapper = newMockedErrorTagMapper()

	// when
	error := sut.validate(nil)

	// then
	assert.Nil(error, "There should not be error returned.")
}
