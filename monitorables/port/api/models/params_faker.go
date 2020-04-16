//+build faker

package models

import (
	"github.com/monitoror/monitoror/internal/pkg/validator"
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	PortParams struct {
		Hostname string `json:"hostname" query:"hostname"`
		Port     int    `json:"port" query:"port"`

		Status coreModels.TileStatus `json:"status" query:"status"`
	}
)

func (p *PortParams) Validate() []validator.Error {
	return validator.Validate(p)
}
