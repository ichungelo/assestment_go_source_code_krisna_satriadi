package utils

import "github.com/go-playground/validator/v10"

// Validate data with go validator
func Validate(dataStruct interface{}) error {
	validateData := validator.New()

	err := validateData.Struct(dataStruct)
	return err
}
