package configuration

import "gopkg.in/alecthomas/kingpin.v2"

type validator struct {
	configuration  *Configuration
	errorTagMapper errorTagMapper
}

func newValidator(configuration *Configuration) *validator {
	errorTagMapper := newValidationErrorMapper()
	return &validator{configuration: configuration, errorTagMapper: errorTagMapper}
}

func (v *validator) validate(*kingpin.Application) error {
	var error error

	if error = v.validateHeaders(); nil != error {
		return error
	}

	if error = v.validateMethods(); nil != error {
		return error
	}

	if error = v.validateWorkersNumber(); nil != error {
		return error
	}

	if error = v.validateBaseURL(); nil != error {
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
		if *v.configuration.workersNumber == 0 {
			return v.errorTagMapper.mapErrorTag(zeroWorkersNumberError)
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
