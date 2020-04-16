package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"

	"github.com/monitoror/monitoror/pkg/humanize"
)

var (
	regexTag = "regex"

	validatorTagMapping = map[string]ErrorID{
		"required": ErrorRequired,
		"gte":      ErrorGTE,
		"gt":       ErrorGT,
		"lte":      ErrorLTE,
		"lt":       ErrorLT,
		"url":      ErrorURL,
		regexTag:   ErrorRegex,
	}
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation(regexTag, ValidateRegex)
}

func Validate(s interface{}) []Error {
	var errors []Error

	if err := validate.Struct(s); err != nil {
		// range over all validate error to bind then into ValidatorError
		for _, err := range err.(validator.ValidationErrors) {
			id, exists := validatorTagMapping[err.Tag()]
			if !exists {
				panic("unsupported validate tag. used tag listed in validatorTagMapping instead.")
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

// ValidateRegex implements validator.Func
func ValidateRegex(fl validator.FieldLevel) bool {
	_, err := regexp.Compile(fl.Field().String())
	return err == nil
}
