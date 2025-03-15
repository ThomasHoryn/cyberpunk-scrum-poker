package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("alphanumspace", validateAlphanumSpace)
	}
}

func validateAlphanumSpace(fl validator.FieldLevel) bool {
	// Use direct regex validation instead of combining validators
	return regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString(fl.Field().String())
}
