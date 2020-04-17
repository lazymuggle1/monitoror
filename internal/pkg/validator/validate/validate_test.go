package validate

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type Params struct {
	URL              *string `validate:"required,url,http"`
	Token            string  `validate:"required"`
	Equal            int     `validate:"eq=0"`
	NotEqual         int     `validate:"ne=1"`
	GreaterThan      int     `validate:"gt=-1"`
	GreaterThanEqual int     `validate:"gte=0"`
	LessThan         int     `validate:"lt=1"`
	LessThanEqual    int     `validate:"lte=0"`
	Omitempty        *int    `validate:"omitempty,gt=0"`
	OneOf            string  `validate:"omitempty,oneof=test"`
	Regex            string  `validate:"regex"`
	Other            string
}

type ErroredTagParams struct {
	Other string `validate:"test"`
}

type UnsupportedTagParams struct {
	Other string `validate:"min=10"`
}

func TestValidate_WithError(t *testing.T) {
	for _, testcase := range []struct {
		params         *Params
		errorID        validator.ErrorID
		errorValue     string
		errorFieldName string
		errorExpected  string
	}{
		{
			params:         &Params{Token: "xxx"},
			errorID:        validator.ErrorRequired,
			errorFieldName: "URL",
		},
		{
			params: &Params{
				URL:   pointer.ToString("http%sexemple.com"),
				Token: "xxx",
			},
			errorID:        validator.ErrorURL,
			errorFieldName: "URL",
			errorValue:     "http%sexemple.com",
		},
		{
			params: &Params{
				URL:   pointer.ToString("ftp://exemple.com"),
				Token: "xxx",
			},
			errorID:        validator.ErrorHTTP,
			errorFieldName: "URL",
			errorValue:     "ftp://exemple.com",
		},
		{
			params: &Params{
				URL: pointer.ToString("http://exemple.com"),
			},
			errorID:        validator.ErrorRequired,
			errorFieldName: "Token",
		},
		{
			params: &Params{
				URL:   pointer.ToString("http://exemple.com"),
				Token: "xxxx",
				Equal: 1000,
			},
			errorID:        validator.ErrorEq,
			errorFieldName: "Equal",
			errorValue:     "1000",
			errorExpected:  "Equal = 0",
		},
		{
			params: &Params{
				URL:      pointer.ToString("http://exemple.com"),
				Token:    "xxxx",
				NotEqual: 1,
			},
			errorID:        validator.ErrorNE,
			errorFieldName: "NotEqual",
			errorValue:     "1",
			errorExpected:  "NotEqual != 1",
		},
		{
			params: &Params{
				URL:         pointer.ToString("http://exemple.com"),
				Token:       "xxxx",
				GreaterThan: -1000,
			},
			errorID:        validator.ErrorGT,
			errorFieldName: "GreaterThan",
			errorValue:     "-1000",
			errorExpected:  "GreaterThan > -1",
		},
		{
			params: &Params{
				URL:              pointer.ToString("http://exemple.com"),
				Token:            "xxxx",
				GreaterThanEqual: -1000,
			},
			errorID:        validator.ErrorGTE,
			errorFieldName: "GreaterThanEqual",
			errorValue:     "-1000",
			errorExpected:  "GreaterThanEqual >= 0",
		},
		{
			params: &Params{
				URL:      pointer.ToString("http://exemple.com"),
				Token:    "xxxx",
				LessThan: 1000,
			},
			errorID:        validator.ErrorLT,
			errorFieldName: "LessThan",
			errorValue:     "1000",
			errorExpected:  "LessThan < 1",
		},
		{
			params: &Params{
				URL:           pointer.ToString("http://exemple.com"),
				Token:         "xxxx",
				LessThanEqual: 1000,
			},
			errorID:        validator.ErrorLTE,
			errorFieldName: "LessThanEqual",
			errorValue:     "1000",
			errorExpected:  "LessThanEqual <= 0",
		},
		{
			params: &Params{
				URL:       pointer.ToString("http://exemple.com"),
				Token:     "xxxx",
				Omitempty: pointer.ToInt(0),
			},
			errorID:        validator.ErrorGT,
			errorFieldName: "Omitempty",
			errorValue:     "0",
			errorExpected:  "Omitempty > 0",
		},
		{
			params: &Params{
				URL:   pointer.ToString("http://exemple.com"),
				Token: "xxxx",
				OneOf: "test2",
			},
			errorID:        validator.ErrorOneOf,
			errorFieldName: "OneOf",
			errorValue:     "test2",
			errorExpected:  "test",
		},
		{
			params: &Params{
				URL:   pointer.ToString("http://exemple.com"),
				Token: "xxxx",
				Regex: "(",
			},
			errorID:        validator.ErrorRegex,
			errorFieldName: "Regex",
			errorValue:     "(",
		},
	} {
		errors := Struct(testcase.params)
		assert.NotEmpty(t, errors)
		castedError := errors[0].(*validateError)
		assert.Equal(t, testcase.errorID, castedError.errorID)
		assert.Equal(t, testcase.errorFieldName, castedError.fieldName)
		if castedError.value != "" {
			assert.Equal(t, testcase.errorValue, castedError.value)
		}
		if castedError.Expected() != "" {
			assert.Equal(t, testcase.errorExpected, castedError.Expected())
		}
	}
}

func TestValidate_Empty(t *testing.T) {
	param := &Params{}
	assert.Len(t, Struct(param), 2)
}

func TestValidate_Panic(t *testing.T) {
	param1 := &UnsupportedTagParams{}
	assert.Panics(t, func() { Struct(param1) })

	param2 := &ErroredTagParams{}
	assert.Panics(t, func() { Struct(param2) })
}

func TestValidate_Success(t *testing.T) {
	param := &Params{
		URL:   pointer.ToString("http://exemple.com"),
		Token: "xxxx",
	}

	assert.Len(t, Struct(param), 0)
}
