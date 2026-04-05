package utils

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

type ErrorDetail struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type IValidator interface {
	ValidateStruct(s interface{}) []*ErrorDetail
}

type customValidator struct {
	validator *validator.Validate
}

var (
	validatorInstance IValidator
	validatorOnce     sync.Once
)

func GetValidator() IValidator {
	validatorOnce.Do(func() {
		validatorInstance = &customValidator{
			validator: validator.New(),
		}
	})
	return validatorInstance
}

func (v *customValidator) ValidateStruct(s interface{}) []*ErrorDetail {
	var errors []*ErrorDetail
	err := v.validator.Struct(s)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range validationErrs {
				var element ErrorDetail
				element.Field = err.Field()
				element.Tag = err.Tag()
				element.Value = err.Param()
				errors = append(errors, &element)
			}
		}
	}
	return errors
}
