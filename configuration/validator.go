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

func (validator *validator) validate(*kingpin.Application) error {
	var error error

	if error = validator.validateHeaders(); nil != error {
		return error
	}

	if error = validator.validateMethods(); nil != error {
		return error
	}

	if error = validator.validateWorkersNumber(); nil != error {
		return error
	}

	if error = validator.validateBaseURL(); nil != error {
		return error
	}

	return nil
}

func (validator *validator) validateHeaders() error {
	if nil != validator.configuration.headers {
		for headerName, headerValue := range *validator.configuration.headers {
			if headerValue == "" {
				return validator.errorTagMapper.mapErrorTag(missingHeaderValueError, headerName)
			}
		}
	}
	return nil
}

func (validator *validator) validateMethods() error {
	if nil != validator.configuration.methods && len(*validator.configuration.methods) > 0 {
		configuredMethods := map[string]bool{}
		for _, method := range *validator.configuration.methods {
			if _, exists := configuredMethods[method]; exists {
				return validator.errorTagMapper.mapErrorTag(repeatedHTTPMethodError, method)
			}
			configuredMethods[method] = true
		}
	}
	return nil
}

func (validator *validator) validateWorkersNumber() error {
	if nil != validator.configuration.workersNumber {
		if *validator.configuration.workersNumber == 0 {
			return validator.errorTagMapper.mapErrorTag(zeroWorkersNumberError)
		}
	}
	return nil
}

func (validator *validator) validateBaseURL() error {
	if nil != validator.configuration.baseURL && nil != *validator.configuration.baseURL && !(**validator.configuration.baseURL).IsAbs() {
		return validator.errorTagMapper.mapErrorTag(relativeBaseURLError, (**validator.configuration.baseURL).String())
	}
	return nil
}
