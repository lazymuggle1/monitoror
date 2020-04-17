package validator

import "fmt"

type (
	Error interface {
		error
		GetErrorID() ErrorID
		SetFieldName(string)
		GetFieldName() string
		GetValue() interface{}
		Expected() string
	}

	ErrorID int
)

// ErrorID enum
const (
	// Default Error
	ErrorDefault ErrorID = iota

	// Validate Error
	ErrorRequired
	ErrorOneOf
	ErrorEq
	ErrorNE
	ErrorGT
	ErrorGTE
	ErrorLT
	ErrorLTE
	ErrorURL
	ErrorNotEmpty
	ErrorHTTP
	ErrorRegex

	// Available Error
)

type DefaultError struct {
	fieldName string
	value     interface{}
	expected  string
}

func NewDefaultError(fieldName string, value interface{}, expected string) Error {
	return &DefaultError{
		fieldName: fieldName,
		value:     value,
		expected:  expected,
	}
}

func (e *DefaultError) Error() string {
	if e.expected != "" {
		return fmt.Sprintf(`Invalid %q field. Must be %s.`, e.fieldName, e.expected)
	}
	return fmt.Sprintf(`Invalid %q field.`, e.fieldName)
}

func (e *DefaultError) GetErrorID() ErrorID   { return ErrorDefault }
func (e *DefaultError) SetFieldName(f string) { e.fieldName = f }
func (e *DefaultError) GetFieldName() string  { return e.fieldName }
func (e *DefaultError) GetValue() interface{} { return e.value }
func (e *DefaultError) Expected() string      { return e.expected }
