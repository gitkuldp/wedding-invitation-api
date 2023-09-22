package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BrideHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Find(ctx echo.Context) error
	List(ctx echo.Context) error
}

type brideHandler struct {
	brideService services.BrideService
}

// Create Bride godoc
//
//	@Summary		Creates a brides type
//	@Description	creates a brides type with given data
//	@Tags			brides
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		models.BrideView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.Bride}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/brides [post]
func (handler brideHandler) Create(ctx echo.Context) error {
	var bride models.BrideView
	err := ctx.Bind(&bride)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(bride)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}
	invite, err := handler.brideService.CreateBride(bride)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

// Detail Bride godoc
//
//	@Summary		Get brides by id
//	@Description	fetches brides against an id
//	@Tags			brides
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		string	true	"Bride ID"
//
//	@Success		200	{object}	api.Response{data=models.Bride}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/brides/{id} [get]
func (handler brideHandler) Find(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "bride not found.")
	}
	bride, err := handler.brideService.FindBride(id)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, bride)
}

// List Bride godoc
//
//	@Summary		Get all brides
//	@Description	fetches all brides from the database
//	@Tags			brides
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	api.Response{data=[]models.Bride}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/brides [get]
func (handler brideHandler) List(ctx echo.Context) error {
	bride, err := handler.brideService.ListBride()
	if err != nil {
		return api.ServerErrorResponse(ctx)
	}
	return api.SuccessResponse(ctx, bride)
}

// Update Bride godoc
//
//	@Summary		Update brides by id
//	@Description	updates brides by id from the database
//	@Tags			brides
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		string				true	"Bride ID"
//
//
//	@Param			request	body		models.BrideView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.Bride}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/brides/{id} [post]
func (handler brideHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "bride not found")
	}
	var bride models.BrideView
	err = ctx.Bind(&bride)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(bride)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	invite, err := handler.brideService.UpdateBride(id, bride)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

func NewBrideHandler(brideService services.BrideService) BrideHandler {
	return brideHandler{
		brideService: brideService,
	}
}
