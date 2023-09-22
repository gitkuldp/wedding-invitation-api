package v1

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AttachmentHandler interface {
	Create(ctx echo.Context) error
	List(ctx echo.Context) error
	Detail(ctx echo.Context) error
}

type attachmentHandler struct {
	attachmentService services.AttachmentService
}

// Create Attachment godoc
//
//	@Summary		List a attachments type
//	@Description	creates a attachments type with given data
//	@Tags			attachments
//	@Accept			multipart/form-data
//	@Produce		json
//
//	@Param			request	formData	models.AttachmentView	true	"request body"
//
//	@Param			images	formData	file					true	"Multiple image files"
//
//	@Success		200		{object}	api.Response{data=models.Attachment}
//	@Failure		400		{object}	api.Response
//	@Failure		404		{object}	api.Response
//	@Failure		500		{object}	api.Response
//	@Router			/attachments [post]
func (handler attachmentHandler) Create(ctx echo.Context) error {
	var attachment models.AttachmentView
	//	@Param	request	formData	models.AttachmentView	true	"request body"

	err := ctx.Bind(&attachment)
	if err != nil {
		return api.BadRequestResponse(ctx, "Bad Request.")
	}
	err = ctx.Validate(attachment)
	if err != nil {
		result := models.TranslateError(err)
		return api.ValidationErrorResponse(ctx, result)
	}

	invite, err := handler.attachmentService.CreateAttachment(ctx, attachment)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, invite)
}

// Detail Attachment godoc
//
//	@Summary		Get attachments by id
//	@Description	fetches attachments against an id
//	@Tags			attachments
//	@Accept			json
//	@Produce		json
//
//	@Param			id	path		string	true	"Attachment ID"
//
//	@Success		200	{object}	api.Response{data=models.Attachment}
//	@Failure		400	{object}	api.Response
//	@Failure		404	{object}	api.Response
//	@Failure		500	{object}	api.Response
//	@Router			/attachments/{id} [get]
func (handler attachmentHandler) Detail(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return api.NotFoundResponse(ctx, "attachment not found.")
	}
	attachment, err := handler.attachmentService.DetailAttachment(id)
	if err != nil {
		return api.ErrorResponse(ctx, err)
	}
	return api.SuccessResponse(ctx, attachment)
}

// List Attachment godoc
//
//	@Summary		List a attachments type
//	@Description	list a attachments type with given data
//	@Tags			attachments
//	@Accept			json
//	@Produce		json
//
//	@Param			is_active	query		bool	false	"Filter Deal Stage by inactive status"
//
//	@Param			pagination	query		bool	false	"State if pagination is required or not"
//	@Param			page		query		int		false	"Enter Page number"
//	@Param			limit		query		int		false	"Enter Page limit"	true	"Multiple image files"
//
//	@Success		200			{object}	api.Response{data=[]models.Attachment}
//	@Failure		400			{object}	api.Response
//	@Failure		404			{object}	api.Response
//	@Failure		500			{object}	api.Response
//	@Router			/attachments [get]
func (handler attachmentHandler) List(ctx echo.Context) error {
	queryParams := ctx.Get("query-params").(models.QueryParams)
	attachments, pagination, err := handler.attachmentService.ListAttachment(ctx, queryParams)
	if err != nil {
		return api.ServerErrorResponse(ctx)
	}
	return models.SuccessResponseWithPagination(ctx, attachments, pagination)
}

func NewAttachmentHandler(attachmentService services.AttachmentService) AttachmentHandler {
	return attachmentHandler{
		attachmentService: attachmentService,
	}
}
