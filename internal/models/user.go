package models

import (
	"time"

	"github.com/gitkuldp/wedding-invitation-api/internal/api"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserLoginAttempt
	ID                 uuid.UUID `json:"id"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	IsActive           bool      `json:"is_active"`
	TokenLastRevokedAt time.Time `json:"-"`
	CreatedAt          time.Time `json:"created_at"`
}

type UserRole struct {
	IdModel
	UserId uuid.UUID     `json:"user_id"`
	User   User          `json:"user"`
	RoleId uuid.UUID     `json:"role_id"`
	Role   RoleReference `json:"role"`
}

func (ur *User) GetUserRole(db *gorm.DB) *uuid.UUID {
	var userRole UserRole
	if err := db.Where("user_id = ?", ur.ID).First(&userRole).Error; err != nil {
		return nil
	}

	return utils.Ref(userRole.RoleId)
}

type UserView struct {
	Phone    string `json:"phone" validate:"required,len=13,number"`
	Email    string `json:"email" validate:"required,lowercase,email,contains=@"`
	Password string `json:"password" validate:"required"`
}

type Invitation struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id,omitempty"`
	User      User
	Title     string `json:"title"`
	Quotes    string `json:"quotes"`
	Moment    string `json:"moment"`
	LoveStory string `json:"lover_story"`
}

type Attachment struct {
	ID        uuid.UUID
	TableName string `json:"table_name"`
	TableID   string `json:"table_id"`
	FileName  string `json:"file_name"`
	FileSize  int    `json:"file_size"`
	FileType  string `json:"file_type"`
	FileURL   string `json:"file_url"`
	IsActive  bool   `json:"is_active"`
	CreatedAt time.Time
}

type AttachmentView struct {
	TableName api.TableName `json:"table_name"`
	TableID   string        `json:"table_id"`
}

type InvitationView struct {
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Quotes    string    `json:"quotes"`
	Moment    string    `json:"moment"`
	LoveStory string    `json:"lover_story"`
}

type Bride struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	User   User

	Name           string `json:"name"`
	InstaID        string `json:"insta_id"`
	Address        string `json:"address"`
	FatherName     string `json:"father_name"`
	MotherName     string `json:"mother_name"`
	ChildTo        string `json:"child_to"`
	AdditionalInfo string `json:"additional_info"`
}

type BrideView struct {
	UserID         uuid.UUID `json:"user_id"`
	Name           string    `json:"name"`
	InstaID        string    `json:"insta_id"`
	Address        string    `json:"address"`
	FatherName     string    `json:"father_name"`
	MotherName     string    `json:"mother_name"`
	ChildTo        string    `json:"child_to"`
	AdditionalInfo string    `json:"additional_info"`
}

type Groom struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	User   User

	Name           string `json:"name"`
	InstaID        string `json:"insta_id"`
	Address        string `json:"address"`
	FatherName     string `json:"father_name"`
	MotherName     string `json:"mother_name"`
	ChildTo        string `json:"child_to"`
	AdditionalInfo string `json:"additional_info"`
}

type GroomView struct {
	UserID         uuid.UUID `json:"user_id"`
	Name           string    `json:"name"`
	InstaID        string    `json:"insta_id"`
	Address        string    `json:"address"`
	FatherName     string    `json:"father_name"`
	MotherName     string    `json:"mother_name"`
	ChildTo        string    `json:"child_to"`
	AdditionalInfo string    `json:"additional_info"`
}

type Speech struct {
	ID       uuid.UUID `json:""`
	FullName string    `json:""`
	Address  string    `json:""`
	Media    []string  `json:""`
	message  string    `json:""`
}

type Confirmation struct {
	ID       uuid.UUID `json:""`
	FullName string    `json:""`
	Message  string    `json:""`
}

type InvitationDate struct {
	InvitationID uuid.UUID  `json:""`
	Invitation   Invitation `json:""`
	ID           uuid.UUID  `json:""`
	EventName    string     `json:""`
	Date         time.Time  `json:""`
	Location     string     `json:""`
	MapUrl       string     `json:""`
}
