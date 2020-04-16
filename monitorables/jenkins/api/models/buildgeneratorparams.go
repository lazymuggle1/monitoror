package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	BuildGeneratorParams struct {
		Job string `json:"job" query:"job" validate:"required"`

		// Using Match / Unmatch filter instead of one filter because Golang's standard regex library doesn't have negative look ahead.
		Match   string `json:"match,omitempty" query:"match" validate:"regex"`
		Unmatch string `json:"unmatch,omitempty" query:"unmatch" validate:"regex"`
	}
)

func (p *BuildGeneratorParams) Validate() []validator.Error {
	return validator.Validate(p)
}
