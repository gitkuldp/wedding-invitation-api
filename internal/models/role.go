package models

import (
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
)

type RoleReference struct {
	ID          uuid.UUID       `json:"id"`
	Name        string          `json:"name" gorm:"unique"`
	Permissions utils.ArrayList `json:"permissions"`
}
