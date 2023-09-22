package api

import (
	"net/http"

	weddinginvitation "github.com/gitkuldp/wedding-invitation-api"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Response struct {
	ErrorMessage    string            `json:"message,omitempty"`
	Data            interface{}       `json:"data,omitempty"`
	Error           bool              `json:"error"`
	ValidationError map[string]string `json:"validation_errors,omitempty"`
}
type ResponseWithPagination struct {
	ErrorMessage    string            `json:"message,omitempty"`
	Data            interface{}       `json:"data,omitempty"`
	MetaData        interface{}       `json:"meta_data,omitempty"`
	Error           bool              `json:"error"`
	ValidationError map[string]string `json:"validation_errors,omitempty"`
}

func BadRequestResponse(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusBadRequest, Response{
		ErrorMessage: message,
		Error:        true,
	})
}

func ValidationErrorResponse(ctx echo.Context, errors map[string]string) error {
	return ctx.JSON(http.StatusBadRequest, Response{
		ErrorMessage:    "Validation Error",
		ValidationError: errors,
		Error:           true,
	})
}

func NotFoundResponse(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusNotFound, Response{
		ErrorMessage: message,
		Error:        true,
	})
}

func ServerErrorResponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusInternalServerError, Response{
		ErrorMessage: "Some bad happened.",
		Error:        true,
	})
}

func SuccessResponse(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, Response{
		Data:  data,
		Error: false,
	})
}

func ErrorResponse(ctx echo.Context, err error) error {
	logrus.Error(err)
	switch err.(type) {
	case weddinginvitation.WeddingInvitationError:
		return BadRequestResponse(ctx, err.Error())
	case weddinginvitation.RecordNotFound:
		return NotFoundResponse(ctx, err.Error())
	case weddinginvitation.InvalidOperation:
		return BadRequestResponse(ctx, err.Error())
	case weddinginvitation.InvalidData:
		return BadRequestResponse(ctx, err.Error())
	case weddinginvitation.DuplicateData:
		return BadRequestResponse(ctx, err.Error())
	default:
		return ServerErrorResponse(ctx)
	}
}
