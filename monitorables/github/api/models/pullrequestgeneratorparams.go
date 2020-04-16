package models

import "github.com/monitoror/monitoror/internal/pkg/validator"

type PullRequestGeneratorParams struct {
	Owner      string `json:"owner" query:"owner" validate:"required"`
	Repository string `json:"repository" query:"repository" validate:"required"`
}

func (p *PullRequestGeneratorParams) Validate() []validator.Error {
	return validator.Validate(p)
}
