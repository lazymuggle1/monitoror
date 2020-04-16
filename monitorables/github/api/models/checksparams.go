//+build !faker

package models

import (
	"fmt"

	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type (
	ChecksParams struct {
		Owner      string `json:"owner" query:"owner" validate:"required"`
		Repository string `json:"repository" query:"repository" validate:"required"`
		Ref        string `json:"ref" query:"ref" validate:"required"`
	}
)

func (p *ChecksParams) Validate() []validator.Error {
	return validator.Validate(p)
}

// Used by cache as identifier
func (p *ChecksParams) String() string {
	return fmt.Sprintf("CHECKS-%s-%s-%s", p.Owner, p.Repository, p.Ref)
}
