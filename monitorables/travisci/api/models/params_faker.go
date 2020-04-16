//+build faker

package models

import (
	"fmt"
	"time"

	uiConfigModels "github.com/monitoror/monitoror/api/config/models"
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	BuildParams struct {
		Owner      string `json:"owner" query:"owner" validate:"required"`
		Repository string `json:"repository" query:"repository" validate:"required"`
		Branch     string `json:"branch" query:"branch" validate:"required"`

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

func (p *BuildParams) Validate(_ *uiConfigModels.ConfigVersion) *uiConfigModels.ConfigError {
	if p.Owner == "" {
		return &uiConfigModels.ConfigError{
			ID:      uiConfigModels.ConfigErrorMissingRequiredField,
			Message: fmt.Sprintf(`Required "owner" field is missing.`),
			Data:    uiConfigModels.ConfigErrorData{FieldName: "owner"},
		}
	}

	if p.Repository == "" {
		return &uiConfigModels.ConfigError{
			ID:      uiConfigModels.ConfigErrorMissingRequiredField,
			Message: fmt.Sprintf(`Required "repository" field is missing.`),
			Data:    uiConfigModels.ConfigErrorData{FieldName: "repository"},
		}
	}

	if p.Branch == "" {
		return &uiConfigModels.ConfigError{
			ID:      uiConfigModels.ConfigErrorMissingRequiredField,
			Message: fmt.Sprintf(`Required "branch" field is missing.`),
			Data:    uiConfigModels.ConfigErrorData{FieldName: "branch"},
		}
	}

	return nil
}
