//+build faker

package models

import (
	"github.com/monitoror/monitoror/internal/pkg/validator"
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	PingParams struct {
		Hostname string `json:"hostname" query:"hostname" validate:"required"`

		Status      coreModels.TileStatus `json:"status" query:"status"`
		ValueValues []string              `json:"valueValues" query:"valueValues"`
	}
)

func (p *PingParams) Validate() []validator.Error {
	return validator.Validate(p)
}
