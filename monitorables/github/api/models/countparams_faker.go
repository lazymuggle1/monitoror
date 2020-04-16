//+build faker

package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	CountParams struct {
		Query string `json:"query" query:"query" validate:"required"`

		ValueValues []string `json:"valueValues" query:"valueValues"`
	}
)

func (p *ChecksParams) Validate() []validator.Error {
	return validator.Validate(p)
}
