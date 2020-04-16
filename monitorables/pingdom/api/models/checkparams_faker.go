//+build faker

package models

import (
	"github.com/monitoror/monitoror/internal/pkg/validator"
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	CheckParams struct {
		ID *int `json:"id" query:"id" validate:"required"`

		Status coreModels.TileStatus `json:"status" query:"status"`
	}
)

func (p *CheckParams) Validate() []validator.Error {
	return validator.Validate(p)
}
