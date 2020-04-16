//+build !faker

package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type (
	PingParams struct {
		Hostname string `json:"hostname" query:"hostname" validate:"required"`
	}
)

func (p *PingParams) Validate() []validator.Error {
	return validator.Validate(p)
}
