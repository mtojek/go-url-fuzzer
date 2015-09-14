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

func (configurationValidator *configurationValidator) validate(*kingpin.Application) error {
	var error error

	if error = configurationValidator.validateHeaders(); nil != error {
		return error
	}

	if error = configurationValidator.validateMethods(); nil != error {
		return error
	}

	if error = configurationValidator.validateWorkersNumber(); nil != error {
		return error
	}

	if error = configurationValidator.validateBaseURL(); nil != error {
		return error
	}

	return nil
}

func (configurationValidator *configurationValidator) validateHeaders() error {
	if nil != configurationValidator.configuration.headers {
		for headerName, headerValue := range *configurationValidator.configuration.headers {
			if headerValue == "" {
				return configurationValidator.errorTagMapper.mapErrorTag(missingHeaderValueError, headerName)
			}
		}
	}
	return nil
}

func (configurationValidator *configurationValidator) validateMethods() error {
	if nil != configurationValidator.configuration.methods && len(*configurationValidator.configuration.methods) > 0 {
		configuredMethods := map[string]bool{}
		for _, method := range *configurationValidator.configuration.methods {
			if _, exists := configuredMethods[method]; exists {
				return configurationValidator.errorTagMapper.mapErrorTag(repeatedHTTPMethodError, method)
			}
			configuredMethods[method] = true
		}
	}
	return nil
}

func (configurationValidator *configurationValidator) validateWorkersNumber() error {
	if nil != configurationValidator.configuration.workersNumber {
		if *configurationValidator.configuration.workersNumber == 0 {
			return configurationValidator.errorTagMapper.mapErrorTag(zeroWorkersNumberError)
		}
	}
	return nil
}

func (configurationValidator *configurationValidator) validateBaseURL() error {
	if nil != configurationValidator.configuration.baseURL && nil != *configurationValidator.configuration.baseURL && !(**configurationValidator.configuration.baseURL).IsAbs() {
		return configurationValidator.errorTagMapper.mapErrorTag(relativeBaseURLError, (**configurationValidator.configuration.baseURL).String())
	}
	return nil
}
