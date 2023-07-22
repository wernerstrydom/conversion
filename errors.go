package conversion

import (
	"fmt"
	"reflect"
)

// OverflowError is an error that occurs when a value is too large or too small to be stored in a variable.
type OverflowError struct {
	value, min, max any
}

func NewOverflowError(value any, min any, max any) *OverflowError {
	return &OverflowError{value: value, min: min, max: max}
}

// Error returns the error message.
func (e *OverflowError) Error() string {
	return fmt.Sprintf("value %v is outside the range %v to %v", e.value, e.min, e.max)
}

// FormatError is an error that occurs when the format of a value is invalid.
type FormatError struct {
	value  any
	format string
}

func NewFormatError(value any, format string) *FormatError {
	return &FormatError{value: value, format: format}
}

// Error returns the error message.
func (e *FormatError) Error() string {
	return fmt.Sprintf("value %v is not in the correct format. Expected %s", e.value, e.format)
}

// MappingError is an error that occurs when a value cannot be converted to another type.
type MappingError struct {
	value any
	from  reflect.Type
	to    reflect.Type
}

func NewMappingError(value any, from reflect.Type, to reflect.Type) *MappingError {
	return &MappingError{value: value, from: from, to: to}
}

// Error returns the error message.
func (e *MappingError) Error() string {
	return fmt.Sprintf("value %v of type %s cannot be converted to type %s", e.value, e.from.String(), e.to.String())
}

// ArgumentNilError is an error that occurs when an argument is nil.
type ArgumentNilError struct {
	ArgumentName string
}

func NewArgumentNilError(argumentName string) *ArgumentNilError {
	return &ArgumentNilError{ArgumentName: argumentName}
}

// Error returns the error message.
func (e *ArgumentNilError) Error() string {
	return fmt.Sprintf("argument %s cannot be nil", e.ArgumentName)
}

// ArgumentError is an error that occurs when an argument is invalid.
type ArgumentError struct {
	ArgumentName string
	Message      string
}

func NewArgumentError(argumentName string, message string) *ArgumentError {
	return &ArgumentError{ArgumentName: argumentName, Message: message}
}

// Error returns the error message.
func (e *ArgumentError) Error() string {
	return fmt.Sprintf("argument %s is invalid: %s", e.ArgumentName, e.Message)
}
