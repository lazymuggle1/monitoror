//go:generate mockery -name ParamsValidator -output ../mocks

package models

import (
	"github.com/monitoror/monitoror/internal/pkg/validator"
)

type ParamsValidator interface {
	Validate() []validator.Error
}
