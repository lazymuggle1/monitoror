package validate

import (
	"regexp"
	"strings"

	pkgValidator "github.com/monitoror/monitoror/internal/pkg/validator"
	"github.com/monitoror/monitoror/pkg/humanize"

	"github.com/go-playground/validator/v10"
)

const (
	regexTag    = "regex"
	httpTag     = "http"
	notEmptytag = "notempty"
)

var (
	validateTagMapping = map[string]pkgValidator.ErrorID{
		"required":  pkgValidator.ErrorRequired,
		"eq":        pkgValidator.ErrorEq,
		"ne":        pkgValidator.ErrorNE,
		"oneof":     pkgValidator.ErrorOneOf,
		"gte":       pkgValidator.ErrorGTE,
		"gt":        pkgValidator.ErrorGT,
		"lte":       pkgValidator.ErrorLTE,
		"lt":        pkgValidator.ErrorLT,
		"url":       pkgValidator.ErrorURL,
		notEmptytag: pkgValidator.ErrorNotEmpty,
		httpTag:     pkgValidator.ErrorHTTP,
		regexTag:    pkgValidator.ErrorRegex,
	}
)

// use a single instance of Struct, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation(notEmptytag, validateNotEmpty)
	_ = validate.RegisterValidation(httpTag, validateHTTP)
	_ = validate.RegisterValidation(regexTag, validateRegex)
}

func Struct(s interface{}) []pkgValidator.Error {
	var errors []pkgValidator.Error

	if err := validate.Struct(s); err != nil {
		// range over all validate validateError to bind then into ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			id, exists := validateTagMapping[err.Tag()]
			if !exists {
				panic("unsupported validate tag. use a tag listed in validateTagMapping instead.")
			}

			e := validateError{
				errorID:   id,
				fieldName: err.Field(),
				value:     humanize.Interface(err.Value()),
				tagParam:  err.Param(),
			}

			errors = append(errors, &e)
		}
	}

	return errors
}

// validateRegex implements validator.Func
func validateRegex(fl validator.FieldLevel) bool {
	_, err := regexp.Compile(fl.Field().String())
	return err == nil
}

// validateHTTP implements validator.Func
func validateHTTP(fl validator.FieldLevel) bool {
	return strings.HasPrefix(fl.Field().String(), "http://") || strings.HasPrefix(fl.Field().String(), "https://")
}

// validateHTTP implements validator.Func
func validateNotEmpty(fl validator.FieldLevel) bool {
	return fl.Field().Len() != 0
}
