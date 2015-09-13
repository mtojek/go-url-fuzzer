package configuration

import (
	"errors"

	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

type configurationValidator struct {
	configuration *Configuration
}

func newConfigurationValidator(configuration *Configuration) *configurationValidator {
	return &configurationValidator{configuration: configuration}
}

func (this *configurationValidator) validate(*kingpin.Application) error {
	var error error

	if error = this.validateHeaders(); nil != error {
		return error
	}

	if error = this.validateMethods(); nil != error {
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
				return errors.New(fmt.Sprintf("Missing header value for header \"%v\".", headerName))
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
				return errors.New(fmt.Sprintf("HTTP methods must not repeat themselves, repeated: \"%v\".", method))
			} else {
				configuredMethods[method] = true
			}
		}
	}
	return nil
}

func (this *configurationValidator) validateBaseURL() error {
	if nil != this.configuration.baseURL && nil != *this.configuration.baseURL && !(**this.configuration.baseURL).IsAbs() {
		return errors.New(fmt.Sprintf("The base URL must be absolute, given: \"%v\".", (**this.configuration.baseURL).String()))
	}
	return nil
}

func (this *configurationValidator) validateWorkersNumber() error {
	if nil != this.configuration.workersNumber {
		if *this.configuration.workersNumber == 0 {
			return errors.New("There must be at least one worker.")
		}
	}
	return nil
}
