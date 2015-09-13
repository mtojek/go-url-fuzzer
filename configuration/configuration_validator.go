package configuration

import "gopkg.in/alecthomas/kingpin.v2"

type configurationValidator struct {
	configuration  *Configuration
	errorTagMapper errorTagMapper
}

func newConfigurationValidator(configuration *Configuration) *configurationValidator {
	errorTagMapper := newConfigurationValidationErrorMapper()
	return &configurationValidator{configuration: configuration, errorTagMapper: errorTagMapper}
}

func (this *configurationValidator) validate(*kingpin.Application) error {
	var error error

	if error = this.validateHeaders(); nil != error {
		return error
	}

	if error = this.validateMethods(); nil != error {
		return error
	}

	if error = this.validateWorkersNumber(); nil != error {
		return error
	}

	if error = this.validateBaseURL(); nil != error {
		return error
	}

	return nil
}

func (this *configurationValidator) validateHeaders() error {
	if nil != this.configuration.headers {
		for headerName, headerValue := range *this.configuration.headers {
			if headerValue == "" {
				return this.errorTagMapper.mapErrorTag(missingHeaderValueError, headerName)
			}
		}
	}
	return nil
}

func (this *configurationValidator) validateMethods() error {
	if nil != this.configuration.methods && len(*this.configuration.methods) > 0 {
		configuredMethods := map[string]bool{}
		for _, method := range *this.configuration.methods {
			if _, exists := configuredMethods[method]; exists {
				return this.errorTagMapper.mapErrorTag(repeatedHttpMethodError, method)
			} else {
				configuredMethods[method] = true
			}
		}
	}
	return nil
}

func (this *configurationValidator) validateWorkersNumber() error {
	if nil != this.configuration.workersNumber {
		if *this.configuration.workersNumber == 0 {
			return this.errorTagMapper.mapErrorTag(zeroWorkersNumberError)
		}
	}
	return nil
}

func (this *configurationValidator) validateBaseURL() error {
	if nil != this.configuration.baseURL && nil != *this.configuration.baseURL && !(**this.configuration.baseURL).IsAbs() {
		return this.errorTagMapper.mapErrorTag(relativeBaseUrlError, (**this.configuration.baseURL).String())
	}
	return nil
}
