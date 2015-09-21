package configuration

import (
	"errors"
	"fmt"
)

const (
	missingHeaderValueError = iota
	repeatedHTTPMethodError
	relativeBaseURLError
	unableToConnectToHostBaseURLError
	zeroWorkersNumberError
	tooManyWorkersError
	invalidHTTPErrorCodeError
)

const (
	missingHeaderValueErrorMessage           = "Missing header value for header \"%v\""
	repeatedHTTPMethodErrorMessage           = "HTTP methods must not repeat themselves, repeated: \"%v\""
	relativeBaseURLErrorMessage              = "The base URL must be absolute, given: \"%v\""
	unableToConnectToHostBaseURLErrorMessage = "Unable to connect to the host: \"%v\", error: %v"
	zeroWorkersNumberErrorMessage            = "There must be at least one worker"
	tooManyWorkersErrorMessage               = "Too many workers defined (upto 255)"
	invalidHTTPErrorCodeErrorMessage         = "Invalid HTTP error code given: \"%v\""
	unknownErrorMessage                      = "Unknown error occurred"
)

type validationErrorMapper struct {
	validationErrorMappings map[int]string
}

func newValidationErrorMapper() *validationErrorMapper {
	validationErrorMappings := map[int]string{
		missingHeaderValueError:           missingHeaderValueErrorMessage,
		repeatedHTTPMethodError:           repeatedHTTPMethodErrorMessage,
		relativeBaseURLError:              relativeBaseURLErrorMessage,
		unableToConnectToHostBaseURLError: unableToConnectToHostBaseURLErrorMessage,
		zeroWorkersNumberError:            zeroWorkersNumberErrorMessage,
		tooManyWorkersError:               tooManyWorkersErrorMessage,
		invalidHTTPErrorCodeError:         invalidHTTPErrorCodeErrorMessage,
	}
	return &validationErrorMapper{validationErrorMappings}
}

func (v *validationErrorMapper) mapErrorTag(tag int, values ...interface{}) error {
	if errorMessage, exists := v.validationErrorMappings[tag]; exists {
		return fmt.Errorf(errorMessage, values...)
	}
	return errors.New(unknownErrorMessage)
}
