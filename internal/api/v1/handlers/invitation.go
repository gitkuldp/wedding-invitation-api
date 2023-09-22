package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type InvitationHandler interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Find(ctx echo.Context) error
	List(ctx echo.Context) error
}

type invitationHandler struct {
	invitationService services.InvitationService
}

// Create Invitation godoc
//
//	@Summary		Creates a invitations type
//	@Description	creates a invitations type with given data
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body		models.InvitationView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.Invitation}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/invitations [post]
func (handler invitationHandler) Create(ctx echo.Context) error {
	var invitation models.InvitationView
	err := ctx.Bind(&invitation)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(invitation)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}
	invite, err := handler.invitationService.CreateInvitation(invitation)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

// Detail Invitation godoc
//
//	@Summary		Get invitations by id
//	@Description	fetches invitations against an id
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		string	true	"Invitation ID"
//
//	@Success		200	{object}	api.Response{data=models.Invitation}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/invitations/{id} [get]
func (handler invitationHandler) Find(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "invitation not found.")
	}
	invitation, err := handler.invitationService.FindInvitation(id)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invitation)
}

// List Invitation godoc
//
//	@Summary		Get all invitations
//	@Description	fetches all invitations from the database
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	api.Response{data=[]models.Invitation}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/invitations [get]
func (handler invitationHandler) List(ctx echo.Context) error {
	invitation, err := handler.invitationService.ListInvitation()
	if err != nil {
		return api.ServerErrorResponse(ctx)
	}
	return api.SuccessResponse(ctx, invitation)
}

// Update Invitation godoc
//
//	@Summary		Update invitations by id
//	@Description	updates invitations by id from the database
//	@Tags			invitations
//	@Accept			json
//	@Produce		json
//
//	@Param			id		path		string					true	"Invitation ID"
//
//
//	@Param			request	body		models.InvitationView	true	"request body"
//
//	@Success		200		{object}	api.Response{data=models.Invitation}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/invitations/{id} [post]
func (handler invitationHandler) Update(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "invitation not found")
	}
	var invitation models.InvitationView
	err = ctx.Bind(&invitation)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(invitation)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	invite, err := handler.invitationService.UpdateInvitation(id, invitation)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

func NewInvitationHandler(invitationService services.InvitationService) InvitationHandler {
	return invitationHandler{
		invitationService: invitationService,
	}
}
