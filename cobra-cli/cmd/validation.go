package cmd

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// usage of validate
//   ref: https://qiita.com/iktakahiro/items/2e240147ca3188948a17
func validateShowOpts(p interface{}) error {
	validate := validator.New()
	errs := validate.Struct(p)
	return extractValidationErrors(errs)
}

func extractValidationErrors(err error) error {
	if err != nil {
		var errorText []string
		for _, err := range err.(validator.ValidationErrors) {
			errorText = append(errorText, validationErrorToText(err))
		}
		return fmt.Errorf("Parameter error: %s", strings.Join(errorText, "\n"))
	}
	return nil
}

func validationErrorToText(e validator.FieldError) string {
	f := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", f)
	case "max":
		return fmt.Sprintf("%s cannot be greater than %s", f, e.Param())
	case "min":
		return fmt.Sprintf("%s must be greater than %s", f, e.Param())
	}
	return fmt.Sprintf("%s is not valid %s", e.Field(), e.Value())
}
