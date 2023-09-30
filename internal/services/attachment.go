package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	weddinginvitation "github.com/gitkuldp/wedding-invitation-api"
	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/api/helper"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type attachmentService struct {
	service
}

// DetailAttachment implements AttachmentService
func (as *attachmentService) DetailAttachment(id uuid.UUID) (*models.Attachment, error) {

	var attach models.Attachment
	if err := as.db.First(&attach).Where("id = ?", id).Error; err != nil {
		return nil, as.translate(err)
	}
	return &attach, nil
}

// ListAttachment implements AttachmentService
func (as *attachmentService) ListAttachment(ctx echo.Context, params models.QueryParams) ([]*models.Attachment, *models.Pagination, error) {
	pagination := utils.ParsePagination(ctx)
	db := ApplyQueryParams(as.db, params)

	var attachments []*models.Attachment
	err := db.Find(&attachments).Error
	if err != nil {
		return nil, nil, err
	}
	meta, err := Paginate(db, pagination, len(attachments), models.Attachment{})
	if err != nil {
		return nil, nil, err
	}
	return attachments, meta, nil
}

type AttachmentService interface {
	CreateAttachment(ctx echo.Context, attachment models.AttachmentView) ([]*models.Attachment, error)
	ListAttachment(ctx echo.Context, params models.QueryParams) ([]*models.Attachment, *models.Pagination, error)
	DetailAttachment(id uuid.UUID) (*models.Attachment, error)
}

func NewAttachmentService(db *gorm.DB) AttachmentService {
	return &attachmentService{
		service{
			db: db,
		},
	}
}

func (as *attachmentService) CreateAttachment(ctx echo.Context, attachment models.AttachmentView) ([]*models.Attachment, error) {
	// Get the uploaded files
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}

	myMap := form.Value
	if values, ok := myMap["table_name"]; ok {
		if len(values) > 0 {
			attachment.TableName = api.TableName(values[0])
		}
	}
	if values, ok := myMap["table_id"]; ok {
		if len(values) > 0 {
			attachment.TableID = values[0]
		}
	}

	tableName, isValid := helper.MapTableName(string(attachment.TableName))

	if !isValid {
		return nil, weddinginvitation.NewWeddingInvitationError("Table not Found")
	}

	var attachments []*models.Attachment

	files := form.File["images"]

	// Create the folder path
	folderPath := fmt.Sprintf("./uploads/%s", string(attachment.TableName))
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0700)
	}

	// Process each uploaded file
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()
		// Read the file content
		byteContent, err := ioutil.ReadAll(src)
		if err != nil {
			return nil, err
		}
		attachmentID := uuid.NewString()
		parsedUUID, err := uuid.Parse(attachmentID)
		if err != nil {

			return nil, as.translate(err)
		}
		fileName := fmt.Sprintf(attachmentID + "-" + utils.SanitizeString(file.Filename))
		fileURL := fmt.Sprintf(string(attachment.TableName) + "/" + fileName)
		filePath := filepath.Join(folderPath, fileName)

		// Write the file to the server
		err = ioutil.WriteFile(filePath, byteContent, 0644)
		if err != nil {
			return nil, err
		}

		// Create the attachment object
		attach := &models.Attachment{
			ID:        parsedUUID,
			TableName: tableName,
			TableID:   attachment.TableID,
			FileName:  utils.SanitizeString(file.Filename),
			FileURL:   fileURL,
			FileType:  file.Header.Get("Content-Type"),
			FileSize:  int(file.Size),
			IsActive:  true,
			CreatedAt: time.Now(),
		}

		attachments = append(attachments, attach)
	}
	err = as.Create(attachments)
	if err != nil {
		return nil, as.translate(err)
	}

	return attachments, nil
}
