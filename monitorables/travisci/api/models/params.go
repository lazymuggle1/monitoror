//+build !faker

package models

import (
	"fmt"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	BuildParams struct {
		Owner      string `json:"owner" query:"owner" validate:"required"`
		Repository string `json:"repository" query:"repository" validate:"required"`
		Branch     string `json:"branch" query:"branch" validate:"required"`
	}
)

func (p *BuildParams) Validate() []validator.Error {
	return validator.Validate(p)
}

// Used by cache as identifier
func (p *BuildParams) String() string {
	return fmt.Sprintf("BUILD-%s-%s-%s", p.Owner, p.Repository, p.Branch)
}
