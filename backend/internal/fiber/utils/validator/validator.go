package route_validator

import (
	"gopkg.in/go-playground/validator.v9"
)

type ValidationError struct {
	Field string
	Tag   string
	Value string
}

var validate = validator.New()

type routeValidator[T any] struct {
}

func Validate[T any](t *T) []ValidationError {
	var errors []ValidationError
	err := validate.Struct(t)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			element := ValidationError{
				Field: e.Field(),
				Tag:   e.Tag(),
				Value: e.Param(),
			}
			errors = append(errors, element)
		}
	}
	return errors
}
