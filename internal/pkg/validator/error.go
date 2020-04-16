package validator

import "fmt"

type (
	// Error is validator error struct
	Error struct {
		ErrorID   ErrorID
		FieldName string
		Value     string

		// expected Only used in case of ErrorDefault
		expected string

		// tagParam used in Expected
		tagParam string
	}

	// ErrorID
	ErrorID int
)

func NewError(fieldName string, value string, expected string) *Error {
	return &Error{
		FieldName: fieldName,
		Value:     value,
		expected:  expected,
	}
}

const (
	// ErrorID enum
	ErrorDefault ErrorID = iota
	ErrorRequired
	ErrorOneOf
	ErrorEq
	ErrorNE
	ErrorGT
	ErrorGTE
	ErrorLT
	ErrorLTE
	ErrorURL
	ErrorHTTP
	ErrorRegex
)

func (e Error) Error() string {
	switch e.ErrorID {
	case ErrorDefault:
		if e.expected != "" {
			return fmt.Sprintf(`Invalid %q field. Must be %s.`, e.FieldName, e.expected)
		}
		return fmt.Sprintf(`Invalid %q field.`, e.FieldName)
	case ErrorRequired:
		return fmt.Sprintf(`Required %q field is missing.`, e.FieldName)
	case ErrorOneOf:
		return fmt.Sprintf(`Invalid %q field. Must be one of [%s].`, e.FieldName, e.tagParam)
	case ErrorEq:
		return fmt.Sprintf(`Invalid %q field. Must be equal to %s.`, e.FieldName, e.tagParam)
	case ErrorNE:
		return fmt.Sprintf(`Invalid %q field. Must be not equal to %s.`, e.FieldName, e.tagParam)
	case ErrorGT:
		return fmt.Sprintf(`Invalid %q field. Must be greater than %s.`, e.FieldName, e.tagParam)
	case ErrorGTE:
		return fmt.Sprintf(`Invalid %q field. Must be greater or equal to %s.`, e.FieldName, e.tagParam)
	case ErrorLT:
		return fmt.Sprintf(`Invalid %q field. Must be lower than %s.`, e.FieldName, e.tagParam)
	case ErrorLTE:
		return fmt.Sprintf(`Invalid %q field. Must be lower or equal to %s.`, e.FieldName, e.tagParam)
	case ErrorURL:
		return fmt.Sprintf(`Invalid %q field. Must be a valid URL.`, e.FieldName)
	case ErrorHTTP:
		return fmt.Sprintf(`Invalid %q field. Must be start with "http://" or "https://".`, e.FieldName)
	case ErrorRegex:
		return fmt.Sprintf(`Invalid %q field. Must be a valid golang regex.`, e.FieldName)
	default:
		return ""
	}
}

func (e Error) Expected() string {
	switch e.ErrorID {
	case ErrorDefault:
		return e.expected
	case ErrorOneOf:
		return e.tagParam
	case ErrorEq:
		return fmt.Sprintf(`%s = %s`, e.FieldName, e.tagParam)
	case ErrorNE:
		return fmt.Sprintf(`%s != %s`, e.FieldName, e.tagParam)
	case ErrorGT:
		return fmt.Sprintf(`%s > %s`, e.FieldName, e.tagParam)
	case ErrorGTE:
		return fmt.Sprintf(`%s >= %s`, e.FieldName, e.tagParam)
	case ErrorLT:
		return fmt.Sprintf(`%s < %s`, e.FieldName, e.tagParam)
	case ErrorLTE:
		return fmt.Sprintf(`%s <= %s`, e.FieldName, e.tagParam)
	default:
		return ""
	}
}
