package services

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type groomService struct {
	service
}

// ListGroom implements GroomService
func (gs *groomService) ListGroom() ([]*models.Groom, error) {
	var groom []*models.Groom
	err := gs.db.Find(&groom).Error
	if err != nil {
		return nil, err
	}
	return groom, nil
}

func (gs groomService) fromView(view models.GroomView, ctx echo.Context) *models.Groom {
	return &models.Groom{
		ID: uuid.New(),
		// Photos:         view.Photos,
		Name:           view.Name,
		InstaID:        view.InstaID,
		Address:        view.Address,
		FatherName:     view.FatherName,
		MotherName:     view.MotherName,
		ChildTo:        view.ChildTo,
		UserID:         utils.Deref(GetUserId(ctx)),
		AdditionalInfo: view.AdditionalInfo,
	}
}

// CreateGroom implements GroomService
func (gs *groomService) CreateGroom(view models.GroomView, ctx echo.Context) (*models.Groom, error) {
	groom := gs.fromView(view, ctx)
	err := gs.Create(groom)
	if err != nil {
		return nil, gs.translate(err)
	}
	return groom, err
}

// FindGroom implements GroomService
func (gs *groomService) FindGroom(id uuid.UUID) (*models.Groom, error) {
	var groom models.Groom
	return &groom, gs.Find(id, &groom)
}

// UpdateGroom implements GroomService
func (gs *groomService) UpdateGroom(id uuid.UUID, view models.GroomView, ctx echo.Context) (uuid.UUID, error) {
	groom := gs.fromView(view, ctx)
	groom.ID = id
	_, err := gs.service.Updates(groom, map[string]interface{}{"name": view.Name, "insta_id": view.InstaID, "address": view.Address, "father_name": view.FatherName, "mother_name": view.MotherName, "child_to": view.ChildTo, "additional_info": view.AdditionalInfo})
	if err != nil {
		return uuid.Nil, gs.translate(err)
	}
	return id, err
}

type GroomService interface {
	CreateGroom(view models.GroomView, ctx echo.Context) (*models.Groom, error)
	UpdateGroom(id uuid.UUID, view models.GroomView, ctx echo.Context) (uuid.UUID, error)
	FindGroom(id uuid.UUID) (*models.Groom, error)
	ListGroom() ([]*models.Groom, error)
}

func NewGroomService(db *gorm.DB) GroomService {
	return &groomService{
		service{
			db: db,
		},
	}
}
