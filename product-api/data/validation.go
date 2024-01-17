package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i)

	if errs == nil {
		return nil
	}

	errors := errs.(validator.ValidationErrors)

	var returnErrs []ValidationError
	for _, err := range errors {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// validateSKU validates the SKU is of valid format
func validateSKU(fl validator.FieldLevel) bool {
	//format of SKU: abc-def-vjhvgj
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches := re.FindAllString(fl.Field().String(), -1)

	return len(matches) == 1
}

// ValidationError wraps the validators FieldError
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}
