//+build faker

package models

import (
	"time"

	"github.com/monitoror/monitoror/internal/pkg/validator"
	"github.com/monitoror/monitoror/models"
)

type (
	BuildParams struct {
		Project    string  `json:"project" query:"project" validate:"required"`
		Definition *int    `json:"definition" query:"definition" validate:"required"`
		Branch     *string `json:"branch,omitempty" query:"branch"`

		AuthorName      string `json:"authorName" query:"authorName"`
		AuthorAvatarURL string `json:"authorAvatarURL" query:"authorAvatarURL"`

		Status            models.TileStatus `json:"status" query:"status"`
		PreviousStatus    models.TileStatus `json:"previousStatus" query:"previousStatus"`
		StartedAt         time.Time         `json:"startedAt" query:"startedAt"`
		FinishedAt        time.Time         `json:"finishedAt" query:"finishedAt"`
		Duration          int64             `json:"duration" query:"duration"`
		EstimatedDuration int64             `json:"estimatedDuration" query:"estimatedDuration"`
	}
)

func (p *BuildParams) Validate() []validator.Error {
	return validator.Validate(p)
}
