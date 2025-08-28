package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Customvalidator using the validator package
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func RegisterValidator(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}
