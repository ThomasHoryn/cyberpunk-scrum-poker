package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidators() {
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("alphanumspacebrackets", ValidateAlphanumSpaceBrackets)
    }
}


func ValidateAlphanumSpaceBrackets(fl validator.FieldLevel) bool {
    // Allows letters, numbers, spaces, hyphens, underscores, and square brackets
    return regexp.MustCompile(`^[a-zA-Z0-9 \-_\[\]]+$`).MatchString(fl.Field().String())
}

