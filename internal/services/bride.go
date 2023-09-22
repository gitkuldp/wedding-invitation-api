package services

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type brideService struct {
	service
}

func (bs brideService) fromView(view models.BrideView) *models.Bride {
	return &models.Bride{
		Name:           view.Name,
		InstaID:        view.InstaID,
		Address:        view.Address,
		FatherName:     view.FatherName,
		MotherName:     view.MotherName,
		ChildTo:        view.ChildTo,
		AdditionalInfo: view.AdditionalInfo,
	}
}

// ListBride implements BrideService
func (bs *brideService) ListBride() ([]*models.Bride, error) {
	var bride []*models.Bride
	err := bs.db.Find(&bride).Error
	if err != nil {
		return nil, err
	}
	return bride, nil
}

// CreateBride implements BrideService
func (bs *brideService) CreateBride(view models.BrideView) (*models.Bride, error) {
	bride := bs.fromView(view)
	err := bs.Create(bride)
	if err != nil {
		return nil, bs.translate(err)
	}
	return bride, err
}

// FindBride implements BrideService
func (bs *brideService) FindBride(id uuid.UUID) (*models.Bride, error) {
	var bride models.Bride
	return &bride, bs.Find(id, &bride)
}

// UpdateBride implements BrideService
func (bs *brideService) UpdateBride(id uuid.UUID, view models.BrideView) (uuid.UUID, error) {
	bride := bs.fromView(view)
	bride.ID = id
	_, err := bs.service.Updates(bride, map[string]interface{}{"name": view.Name, "insta_id": view.InstaID, "address": view.Address, "father_name": view.FatherName, "mother_name": view.MotherName, "child_to": view.ChildTo, "additional_info": view.AdditionalInfo})
	if err != nil {
		return uuid.Nil, bs.translate(err)
	}
	return id, err
}

type BrideService interface {
	CreateBride(view models.BrideView) (*models.Bride, error)
	UpdateBride(id uuid.UUID, view models.BrideView) (uuid.UUID, error)
	FindBride(id uuid.UUID) (*models.Bride, error)
	ListBride() ([]*models.Bride, error)
}

func NewBrideService(db *gorm.DB) BrideService {
	return &brideService{
		service{
			db: db,
		},
	}
}
