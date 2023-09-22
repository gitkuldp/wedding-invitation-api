package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type RegisterHandler interface {
	RegisterUser(ctx echo.Context) error
	UserLogin(ctx echo.Context) error
	RefreshToken(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
}

type registerHandler struct {
	userService services.UserService
}

// Refresh token godoc
//
// @Summary		Refresh token
// @Description	Get new tokens with given data
// @Tags			refresh-token
// @Accept			json
// @Produce		json
//
// @Param			request	body		models.RefreshTokenView	true	"request body"
//
// @Success		200		{object}	api.Response{data=models.Token}
// @Failure		400		{object}	api.Response
// @Failure		404		{object}	api.Response
// @Failure		500		{object}	api.Response
// @Router			/refresh-tokens [post]
func (rh registerHandler) RefreshToken(ctx echo.Context) error {
	var refreshToken models.RefreshTokenView
	err := ctx.Bind(&refreshToken)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(refreshToken)

	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	token, err := rh.userService.RefreshToken(refreshToken)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, token)
}

// User login godoc
//
// @Summary		User login
// @Description	User login with given data
// @Tags			login
// @Accept			json
// @Produce		json
//
// @Param			request	body		models.UserLoginView	true	"request body"
//
// @Success		200		{object}	api.Response{data=models.Token}
// @Failure		400		{object}	api.Response
// @Failure		404		{object}	api.Response
// @Failure		500		{object}	api.Response
// @Router			/login [post]
func (rh registerHandler) UserLogin(ctx echo.Context) error {
	var user models.UserLoginView
	err := ctx.Bind(&user)
	logrus.Info(err)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(user)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	token, err := rh.userService.UserLogin(user)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, token)
}

func NewRegisterHandler(userService services.UserService) RegisterHandler {
	return registerHandler{
		userService: userService,
	}
}

// user registration godoc
//
//	@Summary		user-registration
//	@Description	user-registration with given data
//	@Tags			user-register
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		models.UserView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.User}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/user-registration [post]
func (rh registerHandler) RegisterUser(ctx echo.Context) error {
	var user models.UserView
	err := ctx.Bind(&user)
	logrus.Info(err)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(user)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	emailVerification, err := rh.userService.RegisterUser(user)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, emailVerification)
}

// All GetUsers godoc
//
//	@Summary		Get all get-users
//	@Description	fetches all get-users from the database
//	@Tags			get-users
//
//	@Param			Authorization	header		string	true	"Insert your access token"
//
//	@Success		200				{object}	api.Response{data=[]models.User}
//	@Failure		400				{object}	api.Response
//	@Failure		404				{object}	api.Response
//	@Failure		500				{object}	api.Response
//	@Router			/users [get]
func (rh registerHandler) GetUsers(ctx echo.Context) error {
	var user models.User
	users, err := rh.userService.ListUsers(user)
	if err != nil {
		return api.ServerErrorResponse(ctx)
	}
	return api.SuccessResponse(ctx, users)
}
