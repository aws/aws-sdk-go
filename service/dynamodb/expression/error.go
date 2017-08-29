package expression

import (
	"fmt"
)

// InvalidParameterError blah
type InvalidParameterError struct {
	parameterType string
	function      string
}

func (ipe InvalidParameterError) Error() string {
	return fmt.Sprintf("%s error: invalid parameter: %s", ipe.function, ipe.parameterType)
}

func newInvalidParameterError(function, paramType string) InvalidParameterError {
	return InvalidParameterError{
		parameterType: paramType,
		function:      function,
	}
}

// UnsetParameterError blah
type UnsetParameterError struct {
	parameterType string
	function      string
}

func (upe UnsetParameterError) Error() string {
	return fmt.Sprintf("%s error: unset parameter: %s", upe.function, upe.parameterType)
}

func newUnsetParameterError(function, paramType string) UnsetParameterError {
	return UnsetParameterError{
		parameterType: paramType,
		function:      function,
	}
}
