package configuration

import (
	"crypto/tls"
	"net/http"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

type validator struct {
	configuration  *Configuration
	errorTagMapper errorTagMapper
}

func newValidator(configuration *Configuration) *validator {
	errorTagMapper := newValidationErrorMapper()
	return &validator{configuration: configuration, errorTagMapper: errorTagMapper}
}

func (v *validator) validate(*kingpin.Application) error {
	if error := v.validateOffline(); nil != error {
		return error
	}

	if error := v.validateOnline(); nil != error {
		return error
	}

	return nil
}

func (v *validator) validateOffline() error {
	if error := v.validateHeaders(); nil != error {
		return error
	}

	if error := v.validateMethods(); nil != error {
		return error
	}

	if error := v.validateWorkersNumber(); nil != error {
		return error
	}

	if error := v.validateHTTPErrorCode(); nil != error {
		return error
	}

	if error := v.validateBaseURL(); nil != error {
		return error
	}

	return nil
}

func (v *validator) validateOnline() error {
	if error := v.validateHost(); nil != error {
		return error
	}

	return nil

}

func (v *validator) validateHeaders() error {
	if nil != v.configuration.headers {
		for headerName, headerValue := range *v.configuration.headers {
			if headerValue == "" {
				return v.errorTagMapper.mapErrorTag(missingHeaderValueError, headerName)
			}
		}
	}
	return nil
}

func (v *validator) validateMethods() error {
	if nil != v.configuration.methods && len(*v.configuration.methods) > 0 {
		configuredMethods := map[string]bool{}
		for _, method := range *v.configuration.methods {
			if _, exists := configuredMethods[method]; exists {
				return v.errorTagMapper.mapErrorTag(repeatedHTTPMethodError, method)
			}
			configuredMethods[method] = true
		}
	}
	return nil
}

func (v *validator) validateWorkersNumber() error {
	if nil != v.configuration.workersNumber {
		workersNumber := *v.configuration.workersNumber

		if workersNumber == 0 {
			return v.errorTagMapper.mapErrorTag(zeroWorkersNumberError)
		} else if workersNumber >= (2 << 8) {
			return v.errorTagMapper.mapErrorTag(tooManyWorkersError)
		}
	}
	return nil
}

func (v *validator) validateHTTPErrorCode() error {
	if nil != v.configuration.httpErrorCode {
		httpErrorCode := *v.configuration.httpErrorCode

		if httpErrorCode < 100 || httpErrorCode > 599 {
			return v.errorTagMapper.mapErrorTag(invalidHTTPErrorCodeError, httpErrorCode)
		}
	}
	return nil
}

func (v *validator) validateBaseURL() error {
	if nil != v.configuration.baseURL && nil != *v.configuration.baseURL && !(**v.configuration.baseURL).IsAbs() {
		return v.errorTagMapper.mapErrorTag(relativeBaseURLError, (**v.configuration.baseURL).String())
	}
	return nil
}

func (v *validator) validateHost() error {
	baseURL := *v.configuration.baseURL
	host := baseURL.Host
	scheme := baseURL.Scheme

	timeout := time.Duration(v.configuration.URLResponseTimeout())
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}

	_, error := client.Get(scheme + "://" + host)
	if nil != error {
		return v.errorTagMapper.mapErrorTag(unableToConnectToHostBaseURLError, host, error)
	}

	return nil
}
