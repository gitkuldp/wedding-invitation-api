package services

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type invitationService struct {
	service
}

// AddAttachment implements InvitationService
func (is *invitationService) AddAttachment(id uuid.UUID) (*models.Attachment, error) {

	coverPhotos := &models.Attachment{}
	return coverPhotos, nil
}

// CreateInvitation implements InvitationService
func (is *invitationService) CreateInvitation(view models.InvitationView) (*models.Invitation, error) {
	invitation := is.fromView(view)
	invitation.ID = uuid.New()
	err := is.Create(invitation)
	if err != nil {
		return nil, is.translate(err)
	}
	return invitation, err
}

// FindInvitation implements InvitationService
func (is *invitationService) FindInvitation(id uuid.UUID) (*models.Invitation, error) {
	var invitation models.Invitation
	return &invitation, is.Find(id, &invitation)
}

// ListInvitation implements InvitationService
func (is *invitationService) ListInvitation() ([]*models.Invitation, error) {
	var invitation []*models.Invitation
	err := is.db.Preload("User").Find(&invitation).Error
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

// UpdateInvitation implements InvitationService
func (is *invitationService) UpdateInvitation(id uuid.UUID, view models.InvitationView) (uuid.UUID, error) {
	invitation := is.fromView(view)
	invitation.ID = id
	_, err := is.service.Updates(invitation, map[string]interface{}{"title": view.Title, "moments": view.Moment, "love_story": view.LoveStory})
	if err != nil {
		return uuid.Nil, is.translate(err)
	}
	return id, err
}

func (is invitationService) fromView(view models.InvitationView) *models.Invitation {
	return &models.Invitation{
		ID:    uuid.New(),
		Title: view.Title,
		// CoverPhotos: view.CoverPhotos,
		Quotes:    view.Quotes,
		LoveStory: view.LoveStory,
		Moment:    view.Moment,
		UserID:    view.UserID,
	}
}

type InvitationService interface {
	CreateInvitation(view models.InvitationView) (*models.Invitation, error)
	UpdateInvitation(id uuid.UUID, view models.InvitationView) (uuid.UUID, error)
	FindInvitation(id uuid.UUID) (*models.Invitation, error)
	ListInvitation() ([]*models.Invitation, error)
	AddAttachment(id uuid.UUID) (*models.Attachment, error)
}

func NewInvitationService(db *gorm.DB) InvitationService {
	return &invitationService{
		service{
			db: db,
		},
	}
}
