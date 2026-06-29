package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ErrorDetail represents a single validation error
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationError represents a collection of validation errors
type ValidationError struct {
	Errors []ErrorDetail `json:"errors"`
}

// ParseValidationErrors converts validator errors into meaningful messages
func ParseValidationErrors(err error) ValidationError {
	var response ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			response.Errors = append(response.Errors, ErrorDetail{
				Field:   fe.Field(),
				Message: formatFieldError(fe),
			})
		}
	}

	return response
}

// formatFieldError returns a human-readable message for a validation error
func formatFieldError(fe validator.FieldError) string {
	field := formatFieldName(fe.Field())
	tag := fe.Tag()
	param := fe.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, param)
	case "max":
		return fmt.Sprintf("%s must not exceed %s characters", field, param)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, param)
	case "gte":
		return fmt.Sprintf("%s must be at least %s", field, param)
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, param)
	case "lte":
		return fmt.Sprintf("%s must not exceed %s", field, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", field, strings.ReplaceAll(param, " ", ", "))
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// formatFieldName converts CamelCase field names to readable format
func formatFieldName(field string) string {
	var result strings.Builder
	for i, r := range field {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}
	return result.String()
}
