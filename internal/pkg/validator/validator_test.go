package validator

import (
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
)

type Params struct {
	URL    *string `validate:"required,url"`
	Token  string  `validate:"required"`
	Value1 int     `validate:"gt=-1"`
	Value2 int     `validate:"gte=0"`
	Value3 int     `validate:"lt=1"`
	Value4 int     `validate:"lte=0"`
	Regex  string  `validate:"regex"`
	Other  string
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
		errorID        ErrorID
		errorValue     string
		errorFieldName string
		errorExpected  string
	}{
		{
			params:         &Params{Token: "xxx"},
			errorID:        ErrorRequired,
			errorFieldName: "URL",
		},
		{
			params: &Params{
				URL:   pointer.ToString("http%sexemple.com"),
				Token: "xxx",
			},
			errorID:        ErrorURL,
			errorFieldName: "URL",
			errorValue:     "http%sexemple.com",
		},
		{
			params: &Params{
				URL: pointer.ToString("http://exemple.com"),
			},
			errorID:        ErrorRequired,
			errorFieldName: "Token",
		},
		{
			params: &Params{
				URL:    pointer.ToString("http://exemple.com"),
				Token:  "xxxx",
				Value1: -1000,
			},
			errorID:        ErrorGT,
			errorFieldName: "Value1",
			errorValue:     "-1000",
			errorExpected:  "Value1 > -1",
		},
		{
			params: &Params{
				URL:    pointer.ToString("http://exemple.com"),
				Token:  "xxxx",
				Value2: -1000,
			},
			errorID:        ErrorGTE,
			errorFieldName: "Value2",
			errorValue:     "-1000",
			errorExpected:  "Value2 >= 0",
		},
		{
			params: &Params{
				URL:    pointer.ToString("http://exemple.com"),
				Token:  "xxxx",
				Value3: 1000,
			},
			errorID:        ErrorLT,
			errorFieldName: "Value3",
			errorValue:     "1000",
			errorExpected:  "Value3 < 1",
		},
		{
			params: &Params{
				URL:    pointer.ToString("http://exemple.com"),
				Token:  "xxxx",
				Value4: 1000,
			},
			errorID:        ErrorLTE,
			errorFieldName: "Value4",
			errorValue:     "1000",
			errorExpected:  "Value4 <= 0",
		},
		{
			params: &Params{
				URL:   pointer.ToString("http://exemple.com"),
				Token: "xxxx",
				Regex: "(",
			},
			errorID:        ErrorRegex,
			errorFieldName: "Regex",
			errorValue:     "(",
		},
	} {
		errors := Validate(testcase.params)
		assert.NotEmpty(t, errors)
		assert.Equal(t, testcase.errorID, errors[0].ErrorID)
		assert.Equal(t, testcase.errorFieldName, errors[0].FieldName)
		if errors[0].Value != "" {
			assert.Equal(t, testcase.errorValue, errors[0].Value)
		}
		if errors[0].Expected() != "" {
			assert.Equal(t, testcase.errorExpected, errors[0].Expected())
		}
	}
}

func TestValidate_Empty(t *testing.T) {
	param := &Params{}
	assert.Len(t, Validate(param), 2)
}

func TestValidate_Panic(t *testing.T) {
	param1 := &UnsupportedTagParams{}
	assert.Panics(t, func() { Validate(param1) })

	param2 := &ErroredTagParams{}
	assert.Panics(t, func() { Validate(param2) })
}

func TestValidate_Success(t *testing.T) {
	param := &Params{
		URL:   pointer.ToString("http://exemple.com"),
		Token: "xxxx",
	}

	assert.Len(t, Validate(param), 0)
}
