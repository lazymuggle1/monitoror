//+build !faker

package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	CheckParams struct {
		ID *int `json:"id" query:"id" validate:"required"`
	}
)

func (p *CheckParams) Validate() []validator.Error {
	return validator.Validate(p)
}
