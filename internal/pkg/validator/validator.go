package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/monitoror/monitoror/pkg/humanize"
)

const (
	regexTag = "regex"
	httpTag  = "http"
)

var (
	validatorTagMapping = map[string]ErrorID{
		"required": ErrorRequired,
		"eq":       ErrorEq,
		"ne":       ErrorNE,
		"oneof":    ErrorOneOf,
		"gte":      ErrorGTE,
		"gt":       ErrorGT,
		"lte":      ErrorLTE,
		"lt":       ErrorLT,
		"url":      ErrorURL,
		httpTag:    ErrorHTTP,
		regexTag:   ErrorRegex,
	}
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation(httpTag, validateHTTP)
	_ = validate.RegisterValidation(regexTag, validateRegex)
}

func Validate(s interface{}) []Error {
	var errors []Error

	if err := validate.Struct(s); err != nil {
		// range over all validate error to bind then into ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			id, exists := validatorTagMapping[err.Tag()]
			if !exists {
				panic("unsupported validate tag. use a tag listed in validatorTagMapping instead.")
			}

			e := Error{
				ErrorID:   id,
				FieldName: err.Field(),
				Value:     humanize.Interface(err.Value()),
				tagParam:  err.Param(),
			}

			errors = append(errors, e)
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
