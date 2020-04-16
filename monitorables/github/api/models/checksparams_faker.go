//+build faker

package models

import (
	"fmt"
	"time"

	"github.com/monitoror/monitoror/internal/pkg/validator"
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	ChecksParams struct {
		Owner      string `json:"owner" query:"owner" validate:"required"`
		Repository string `json:"repository" query:"repository" validate:"required"`
		Ref        string `json:"ref" query:"ref" validate:"required"`

		AuthorName      string `json:"authorName" query:"authorName"`
		AuthorAvatarURL string `json:"authorAvatarURL" query:"authorAvatarURL"`

		Status            coreModels.TileStatus `json:"status" query:"status"`
		PreviousStatus    coreModels.TileStatus `json:"previousStatus" query:"previousStatus"`
		StartedAt         time.Time             `json:"startedAt" query:"startedAt"`
		FinishedAt        time.Time             `json:"finishedAt" query:"finishedAt"`
		Duration          int64                 `json:"duration" query:"duration"`
		EstimatedDuration int64                 `json:"estimatedDuration" query:"estimatedDuration"`
	}
)

func (p *ChecksParams) Validate() []validator.Error {
	return validator.Validate(p)
}

// Used by cache as identifier
func (p *ChecksParams) String() string {
	return fmt.Sprintf("CHECKS-%s-%s-%s", p.Owner, p.Repository, p.Ref)
}
