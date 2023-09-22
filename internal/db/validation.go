package db

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type validate struct {
	validator *validator.Validate
}

func (v validate) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func Date(fl validator.FieldLevel) bool {
	_, err := time.Parse("02-01-2006", fl.Field().String())
	return err == nil
}

func NewValidator() echo.Validator {
	val := validator.New()
	_ = val.RegisterValidation("date", Date)

	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return validate{validator: val}
}
