package validator

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestError(t *testing.T) {
	for _, testcase := range []struct {
		err      *Error
		message  string
		expected string
	}{
		{
			err:      NewError("test", "test", "expected"),
			message:  `Invalid "test" field. Must be expected.`,
			expected: `expected`,
		},
		{
			err:      NewError("test", "test", ""),
			message:  `Invalid "test" field.`,
			expected: "",
		},
		{
			err:      &Error{ErrorID: ErrorRequired, FieldName: "test"},
			message:  `Required "test" field is missing.`,
			expected: "",
		},
		{
			err:      &Error{ErrorID: ErrorGT, FieldName: "test", tagParam: "1"},
			message:  `Invalid "test" field. Must be greater than 1.`,
			expected: `test > 1`,
		},
		{
			err:      &Error{ErrorID: ErrorGTE, FieldName: "test", tagParam: "1"},
			message:  `Invalid "test" field. Must be greater or equal to 1.`,
			expected: `test >= 1`,
		},
		{
			err:      &Error{ErrorID: ErrorLT, FieldName: "test", tagParam: "1"},
			message:  `Invalid "test" field. Must be lower than 1.`,
			expected: `test < 1`,
		},
		{
			err:      &Error{ErrorID: ErrorLTE, FieldName: "test", tagParam: "1"},
			message:  `Invalid "test" field. Must be lower or equal to 1.`,
			expected: `test <= 1`,
		},
		{
			err:      &Error{ErrorID: ErrorURL, FieldName: "test"},
			message:  `Invalid "test" field. Must be a valid URL.`,
			expected: "",
		},
		{
			err:      &Error{ErrorID: ErrorRegex, FieldName: "test"},
			message:  `Invalid "test" field. Must be a valid golang regex.`,
			expected: "",
		},
		{
			err:      &Error{ErrorID: 99999},
			message:  "",
			expected: "",
		},
	} {
		assert.Equal(t, testcase.message, testcase.err.Error())
		assert.Equal(t, testcase.expected, testcase.err.Expected())
	}
}
