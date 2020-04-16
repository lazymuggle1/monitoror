package models

import (
	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	CheckGeneratorParams struct {
		Tags   string `json:"tags,omitempty" query:"tags"`
		SortBy string `json:"sortBy,omitempty" query:"sortBy"`
	}
)

func (p *CheckGeneratorParams) Validate() []validator.Error {
	errors := validator.Validate(p)

	if p.SortBy != "" && p.SortBy != "name" {
		errors = append(errors, *validator.NewError("SortBy", p.SortBy, `"name"`))
	}

	return errors
}
