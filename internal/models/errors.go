package models

import (
	"net/http"
	"strings"

	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Required!"
	case "required_if":
		return "Required!"
	case "min":
		if fe.Type().String() == "string" {
			return "Value must be at least " + fe.Param() + " characters long."
		}
		return "At least " + fe.Param() + " item required."
	case "max":
		return "Value must be at max " + fe.Param() + " characters long."
	case "gte":
		return "Value must be >= " + fe.Param()
	case "lte":
		return "Value must be <= " + fe.Param()
	case "len":
		return "Value must have length = " + fe.Param()
	case "numeric":
		return "Value must be numeric data"
	case "oneof":
		return "Invalid!"
	case "hexcolor":
		return "Value must be hexcolor "
	case "date":
		return "invalid date"
	case "email":
		return "invalid email format"
	case "accounting_date":
		return "Date can not be in the future"
	}

	return fe.Error() // default error
}

func TranslateError(err error) map[string]string {
	var result = map[string]string{}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			key := err.Field()
			count := strings.Count(err.Namespace(), ".")

			if count > 1 {
				index := strings.Index(err.Namespace(), ".")
				key = err.Namespace()[index+1:]
			}
			result[key] = msgForTag(err)
		}
	}
	return result
}

func SuccessResponseWithPagination(ctx echo.Context, data any, pagination *Pagination) error {
	return ctx.JSON(http.StatusOK, api.ResponseWithPagination{
		Data:     data,
		MetaData: pagination,
		Error:    false,
	})
}
