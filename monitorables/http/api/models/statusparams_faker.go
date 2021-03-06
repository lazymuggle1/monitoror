//+build faker

package models

import (
	uiConfigModels "github.com/monitoror/monitoror/api/config/models"
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	HTTPStatusParams struct {
		URL           string `json:"url" query:"url"`
		StatusCodeMin *int   `json:"statusCodeMin,omitempty" query:"statusCodeMin"`
		StatusCodeMax *int   `json:"statusCodeMax,omitempty" query:"statusCodeMax"`

		Status  coreModels.TileStatus `json:"status" query:"status"`
		Message string                `json:"message" query:"message"`
	}
)

func (p *HTTPStatusParams) Validate(_ *uiConfigModels.ConfigVersion) *uiConfigModels.ConfigError {
	if err := validateURL(p); err != nil {
		return err
	}

	if err := validateStatusCode(p); err != nil {
		return err
	}

	return nil
}

func (p *HTTPStatusParams) GetURL() (url string) { return p.URL }
func (p *HTTPStatusParams) GetStatusCodes() (min int, max int) {
	return getStatusCodesWithDefault(p.StatusCodeMin, p.StatusCodeMax)
}

func (p *HTTPStatusParams) GetStatus() coreModels.TileStatus        { return p.Status }
func (p *HTTPStatusParams) GetMessage() string                      { return p.Message }
func (p *HTTPStatusParams) GetValueValues() []string                { panic("unimplemented") }
func (p *HTTPStatusParams) GetValueUnit() coreModels.TileValuesUnit { panic("unimplemented") }
