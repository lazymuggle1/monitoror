//+build !faker

package models

import (
	"fmt"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	BuildParams struct {
		Job    string `json:"job" query:"job" validate:"required"`
		Branch string `json:"branch,omitempty" query:"branch"`
	}
)

func (p *BuildParams) Validate() []validator.Error {
	return validator.Validate(p)
}

// Used by cache as identifier
func (p *BuildParams) String() string {
	return fmt.Sprintf("BUILD-%s-%s", p.Job, p.Branch)
}
