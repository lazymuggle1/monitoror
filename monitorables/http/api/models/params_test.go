package models

import (
	"regexp"
	"testing"

	uiConfigModels "github.com/monitoror/monitoror/api/config/models"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
)

func TestHTTPParams_GetFormat(t *testing.T) {
	for _, testcase := range []struct {
		params uiConfigModels.ParamsValidator
		valid  bool
	}{
		{&HTTPStatusParams{}, false},
		{&HTTPStatusParams{URL: "example.com"}, false},
		{&HTTPStatusParams{URL: "http%sexample.com"}, false},
		{&HTTPStatusParams{URL: "http://example.com"}, true},
		{&HTTPStatusParams{URL: "http://example.com", StatusCodeMin: pointer.ToInt(300), StatusCodeMax: pointer.ToInt(299)}, false},
		{&HTTPStatusParams{URL: "http://example.com", StatusCodeMin: pointer.ToInt(299), StatusCodeMax: pointer.ToInt(300)}, true},

		{&HTTPRawParams{}, false},
		{&HTTPRawParams{URL: "http://example.com"}, true},
		{&HTTPRawParams{URL: "http://example.com", StatusCodeMin: pointer.ToInt(300), StatusCodeMax: pointer.ToInt(299)}, false},
		{&HTTPRawParams{URL: "http://example.com", StatusCodeMin: pointer.ToInt(299), StatusCodeMax: pointer.ToInt(300)}, true},
		{&HTTPRawParams{URL: "http://example.com", Regex: "("}, false},
		{&HTTPRawParams{URL: "http://example.com", Regex: "(.*)"}, true},

		{&HTTPFormattedParams{}, false},
		{&HTTPFormattedParams{URL: "http://example.com"}, false},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "unknown"}, false},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: ""}, false},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: "."}, false},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: "key"}, true},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: "key", StatusCodeMin: pointer.ToInt(300), StatusCodeMax: pointer.ToInt(299)}, false},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: "key", StatusCodeMin: pointer.ToInt(299), StatusCodeMax: pointer.ToInt(300)}, true},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: "key", Regex: "("}, false},
		{&HTTPFormattedParams{URL: "http://example.com", Format: "JSON", Key: "key", Regex: "(.*)"}, true},
	} {
		errors := testcase.params.Validate()
		if testcase.valid {
			assert.NotEmpty(t, testcase.params.(GenericParamsProvider).GetURL())
			assert.Empty(t, errors)
		} else {
			assert.NotEmpty(t, errors)
		}
	}
}

func TestHTTPParams_GetRegex(t *testing.T) {
	for _, testcase := range []struct {
		params         RegexParamsProvider
		expectedRegex  string
		expectedRegexp *regexp.Regexp
	}{
		{&HTTPRawParams{}, "", nil},
		{&HTTPRawParams{Regex: ""}, "", nil},
		{&HTTPRawParams{Regex: "("}, "(", nil},
		{&HTTPRawParams{Regex: "(.*)"}, "(.*)", regexp.MustCompile("(.*)")},

		{&HTTPFormattedParams{}, "", nil},
		{&HTTPFormattedParams{Regex: ""}, "", nil},
		{&HTTPFormattedParams{Regex: "("}, "(", nil},
		{&HTTPFormattedParams{Regex: "(.*)"}, "(.*)", regexp.MustCompile("(.*)")},
	} {
		assert.Equal(t, testcase.expectedRegex, testcase.params.GetRegex())
		assert.Equal(t, testcase.expectedRegexp, testcase.params.GetRegexp())
	}
}

func TestHTTPFormattedParams_FormattedDataProvider(t *testing.T) {
	for _, testcase := range []struct {
		params         FormattedParamsProvider
		expectedFormat Format
		expectedKey    string
	}{
		{&HTTPFormattedParams{}, "", ""},
		{&HTTPFormattedParams{Format: JSONFormat}, JSONFormat, ""},
		{&HTTPFormattedParams{Format: YAMLFormat, Key: "key"}, YAMLFormat, "key"},
		{&HTTPFormattedParams{Format: XMLFormat, Key: "key"}, XMLFormat, "key"},
	} {
		assert.Equal(t, testcase.expectedFormat, testcase.params.GetFormat())
		assert.Equal(t, testcase.expectedKey, testcase.params.GetKey())
	}
}
