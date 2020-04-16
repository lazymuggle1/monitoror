//+build !faker

package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	PortParams struct {
		Hostname string `json:"hostname" query:"hostname" validate:"required"`
		Port     int    `json:"port" query:"port" validate:"required,gt=0"`
	}
)

func (p *PortParams) Validate() []validator.Error {
	return validator.Validate(p)
}
