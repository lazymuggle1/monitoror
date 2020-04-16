//+build !faker

package models

import (
	"fmt"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	ReleaseParams struct {
		Project    string `json:"project" query:"project" validate:"required"`
		Definition *int   `json:"definition" query:"definition" validate:"required"`
	}
)

func (p *ReleaseParams) Validate() []validator.Error {
	return validator.Validate(p)
}

// Used by cache as identifier
func (p *ReleaseParams) String() string {
	return fmt.Sprintf("RELEASE-%s-%d", p.Project, *p.Definition)
}
