//+build !faker

package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	HTTPStatusParams struct {
		URL           string `json:"url" query:"url" validate:"required,url"`
		StatusCodeMin *int   `json:"statusCodeMin,omitempty" query:"statusCodeMin"`
		StatusCodeMax *int   `json:"statusCodeMax,omitempty" query:"statusCodeMax"`
	}
)

func (p *HTTPStatusParams) Validate() []validator.Error {
	errors := validator.Validate(p)
	errors = append(errors, validateURL(p)...)
	errors = append(errors, validateStatusCode(p)...)
	return errors
}

func (p *HTTPStatusParams) GetURL() (url string) { return p.URL }
func (p *HTTPStatusParams) GetStatusCodes() (min int, max int) {
	return getStatusCodesWithDefault(p.StatusCodeMin, p.StatusCodeMax)
}
