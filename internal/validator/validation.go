package validator

import (
	"github.com/go-playground/validator/v10"
)

var validation = validator.New()

type Errors struct {
	Errors []string
}

func init() {
	validation.RegisterValidation("dateRequired", DateRequired)
	validation.RegisterValidation("gteDate", GteDate)
	validation.RegisterValidation("gtToday", GtToday)
}

func ValidateRequest(req interface{}) Errors {
	var errors Errors

	err := validation.Struct(req)

	if err == nil {
		return errors
	}

	for _, e := range err.(validator.ValidationErrors) {
		errors.Errors = append(errors.Errors, e.Error())
	}

	return errors
}
