//+build !faker

package models

import (
	"fmt"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	BuildParams struct {
		Project    string  `json:"project" query:"project" validate:"required"`
		Definition *int    `json:"definition" query:"definition" validate:"required"`
		Branch     *string `json:"branch,omitempty" query:"branch"`
	}
)

func (p *BuildParams) Validate() []validator.Error {
	return validator.Validate(p)
}

// Used by cache as identifier
func (p *BuildParams) String() string {
	str := fmt.Sprintf("BUILD-%s-%d", p.Project, *p.Definition)

	if p.Branch != nil {
		str = fmt.Sprintf("%s-%s", str, *p.Branch)
	}

	return str
}
