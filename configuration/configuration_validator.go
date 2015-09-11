package configuration

import (
	"errors"

	"gopkg.in/alecthomas/kingpin.v2"
)

type configurationValidator struct {
	configuration *Configuration
}

func newConfigurationValidator(configuration *Configuration) *configurationValidator {
	return &configurationValidator{configuration: configuration}
}

func (this *configurationValidator) validate(*kingpin.Application) error {
	return errors.New("an_error")
}
