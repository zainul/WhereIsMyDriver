package structs

import (
	validator "gopkg.in/go-playground/validator.v9"
)

type (

	// Response is standard response for http
	Response struct {
		Data   interface{} `json:"data"`
		Errors []string    `json:"errors"`
	}
)

// MapValidation mapping the validation of error
func MapValidation(err error) []string {
	var appErrors []string

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// fmt.Println(err)
		}
		for _, errVal := range err.(validator.ValidationErrors) {
			errDesc := mappingStrError(errVal.Tag(), errVal.Param())

			errStr := "Field " + errVal.Field() + " " + errDesc

			appErrors = append(appErrors, errStr)
		}
	}

	return appErrors
}

func mappingStrError(typeErr string, val string) (errStr string) {
	if typeErr == "gte" {
		errStr = "must be greater than or equal " + val + " character"
	} else if typeErr == "lte" {
		errStr = "must be less than or equal " + val + " character"
	} else if typeErr == "email" {
		errStr = "must be have valid email format"
	} else if typeErr == "required" {
		errStr = "is required"
	}

	return
}
