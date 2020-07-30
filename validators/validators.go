package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidateTitle is ...
func ValidateTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "okey")
}
