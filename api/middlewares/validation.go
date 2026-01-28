package middlewares

import (
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func InitValidator() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("containsNumber", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[0-9]`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("containsSpecialCharacter", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`[!@#$%^&*()]`, fl.Field().String())
		return match
	})

	validate.RegisterValidation("validDate", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4})$`, fl.Field().String())
		return match
	})

	return validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	errorMessages := ""

	if err := cv.Validator.Struct(i); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages += getErrorMessage(err.Field(), err)
		}

		return errors.New(errorMessages)
	}
	return nil
}

func getErrorMessage(fieldName string, err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fieldName + " is required"
	case "email":
		return "the email is invalid"
	case "min":
		return "the minimum length of " + fieldName + " is equals " + err.Param()
	case "max":
		return "the maximum length of " + fieldName + " is equals " + err.Param()
	case "containsNumber":
		return "the " + fieldName + " must contains number"
	case "containsSpecialCharacter":
		return "the " + fieldName + " must contains special character"
	case "validDate":
		return "the " + fieldName + " must follows this format: DD-MM-YYYY"
	case "gte":
		return "the " + fieldName + " must be greater than or equal to 1"
	default:
		return "validation error in " + fieldName
	}
}
