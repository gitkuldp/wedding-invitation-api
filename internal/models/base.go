package models

import "github.com/google/uuid"

var AllModels = []interface{}{
	User{},
	UserRole{},
	Invitation{},
	Bride{},
	Groom{},
	Attachment{},
}

type IdModel struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
}
