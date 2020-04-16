//+build !faker

package models

import (
	"regexp"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	HTTPRawParams struct {
		URL           string `json:"url" query:"url" validate:"required,url"`
		Regex         string `json:"regex,omitempty" query:"regex" validate:"regex"`
		StatusCodeMin *int   `json:"statusCodeMin,omitempty" query:"statusCodeMin"`
		StatusCodeMax *int   `json:"statusCodeMax,omitempty" query:"statusCodeMax"`
	}
)

func (p *HTTPRawParams) Validate() []validator.Error {
	errors := validator.Validate(p)
	errors = append(errors, validateURL(p)...)
	errors = append(errors, validateStatusCode(p)...)
	return errors
}

func (p *HTTPRawParams) GetURL() (url string) { return p.URL }
func (p *HTTPRawParams) GetStatusCodes() (min int, max int) {
	return getStatusCodesWithDefault(p.StatusCodeMin, p.StatusCodeMax)
}

func (p *HTTPRawParams) GetRegex() string          { return p.Regex }
func (p *HTTPRawParams) GetRegexp() *regexp.Regexp { return getRegexp(p.GetRegex()) }
