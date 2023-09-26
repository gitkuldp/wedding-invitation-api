package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GroomHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Find(ctx echo.Context) error
	List(ctx echo.Context) error
}

type groomHandler struct {
	groomService services.GroomService
}

// List Groom godoc
//
//	@Summary		Get all grooms
//	@Description	fetches all grooms from the database
//	@Tags			grooms
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	api.Response{data=[]models.Groom}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/grooms [get]
func (handler groomHandler) List(ctx echo.Context) error {
	groom, err := handler.groomService.ListGroom()
	if err != nil {
		return api.ServerErrorResponse(ctx)
	}
	return api.SuccessResponse(ctx, groom)
}

// Create Groom godoc
//
//	@Summary		Creates a grooms type
//	@Description	creates a grooms type with given data
//	@Tags			grooms
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		models.GroomView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.Groom}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/grooms [post]
func (handler groomHandler) Create(ctx echo.Context) error {
	var groom models.GroomView
	err := ctx.Bind(&groom)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(groom)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}
	invite, err := handler.groomService.CreateGroom(groom, ctx)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

// Detail Groom godoc
//
//	@Summary		Get grooms by id
//	@Description	fetches grooms against an id
//	@Tags			grooms
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		string	true	"Groom ID"
//
//	@Success		200	{object}	api.Response{data=models.Groom}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/grooms/{id} [get]
func (handler groomHandler) Find(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "groom not found.")
	}
	groom, err := handler.groomService.FindGroom(id)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, groom)
}

// Update Groom godoc
//
//	@Summary		Update grooms by id
//	@Description	updates grooms by id from the database
//	@Tags			grooms
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		string				true	"Groom ID"
//
//
//	@Param			request	body		models.GroomView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.Groom}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/grooms/{id} [post]
func (handler groomHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "groom not found")
	}
	var groom models.GroomView
	err = ctx.Bind(&groom)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(groom)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	invite, err := handler.groomService.UpdateGroom(id, groom, ctx)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

func NewGroomHandler(groomService services.GroomService) GroomHandler {
	return groomHandler{
		groomService: groomService,
	}
}
