package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ThrowError(err error) map[string]string {
	errors := map[string]string{}
	for _, error := range err.(validator.ValidationErrors) {
		var msg string
		switch error.Tag() {
		case "required":
			msg = "This field is required"
		case "email":
			msg = "Invalid email format"
		case "min":
			msg = "Minimum length is " + error.Param()
		default:
			msg = "Invalid input"
		}
		errors[strings.ToLower(error.Field())] = msg
	}

	return errors
}
