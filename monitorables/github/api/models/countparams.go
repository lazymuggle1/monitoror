//+build !faker

package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	CountParams struct {
		Query string `json:"query" query:"query" validate:"required"`
	}
)

func (p *CountParams) Validate() []validator.Error {
	return validator.Validate(p)
}
