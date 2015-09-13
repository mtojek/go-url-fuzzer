package configuration

import (
	"errors"
	"fmt"
)

const (
	missingHeaderValueError = iota
	repeatedHttpMethodError
	relativeBaseUrlError
	zeroWorkersNumberError
)

const (
	missingHeaderValueErrorMessage = "Missing header value for header \"%v\"."
	repeatedHttpMethodErrorMessage = "HTTP methods must not repeat themselves, repeated: \"%v\"."
	relativeBaseUrlErrorMessage    = "The base URL must be absolute, given: \"%v\"."
	zeroWorkersNumberErrorMessage  = "There must be at least one worker."
	unknownErrorMessage            = "Unknown error occurred."
)

type configurationValidationErrorMapper struct {
	validationErrorMappings map[int]string
}

func newConfigurationValidationErrorMapper() *configurationValidationErrorMapper {
	validationErrorMappings := map[int]string{
		missingHeaderValueError: missingHeaderValueErrorMessage,
		repeatedHttpMethodError: repeatedHttpMethodErrorMessage,
		relativeBaseUrlError:    relativeBaseUrlErrorMessage,
		zeroWorkersNumberError:  zeroWorkersNumberErrorMessage,
	}
	return &configurationValidationErrorMapper{validationErrorMappings}
}

func (this *configurationValidationErrorMapper) mapErrorTag(tag int, values ...interface{}) error {
	if errorMessage, exists := this.validationErrorMappings[tag]; exists {
		return errors.New(fmt.Sprintf(errorMessage, values))
	} else {
		return errors.New(unknownErrorMessage)
	}
}
