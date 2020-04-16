//+build !faker

package models

import (
	"regexp"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	HTTPFormattedParams struct {
		URL           string `json:"url" query:"url" validate:"required,url"`
		Format        Format `json:"format" query:"format" validate:"required"`
		Key           string `json:"key" query:"key" validate:"required"`
		Regex         string `json:"regex,omitempty" query:"regex" validate:"regex"`
		StatusCodeMin *int   `json:"statusCodeMin,omitempty" query:"statusCodeMin"`
		StatusCodeMax *int   `json:"statusCodeMax,omitempty" query:"statusCodeMax"`
	}
)

func (p *HTTPFormattedParams) Validate() []validator.Error {
	errors := validator.Validate(p)
	errors = append(errors, validateURL(p)...)
	errors = append(errors, validateStatusCode(p)...)
	errors = append(errors, validateFormat(p)...)
	errors = append(errors, validateKey(p)...)
	return errors
}

func (p *HTTPFormattedParams) GetURL() (url string) { return p.URL }
func (p *HTTPFormattedParams) GetStatusCodes() (min int, max int) {
	return getStatusCodesWithDefault(p.StatusCodeMin, p.StatusCodeMax)
}

func (p *HTTPFormattedParams) GetRegex() string          { return p.Regex }
func (p *HTTPFormattedParams) GetRegexp() *regexp.Regexp { return getRegexp(p.GetRegex()) }

func (p *HTTPFormattedParams) GetKey() string    { return p.Key }
func (p *HTTPFormattedParams) GetFormat() Format { return p.Format }
