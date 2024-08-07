package validate

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateTodoStruct(todo interface{}) []string {
	var validationErrors []string
	err := validate.Struct(todo)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, err.Field()+" is required")
		}
	}
	return validationErrors
}
