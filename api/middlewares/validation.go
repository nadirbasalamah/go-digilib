package middlewares

import "github.com/go-playground/validator/v10"

func InitValidator() *validator.Validate {
	validate := validator.New()
	return validate
}
