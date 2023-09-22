package services

import (
	"errors"

	weddinginvitation "github.com/LamichhaneBibek/wedding-invitation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func (s service) Create(obj interface{}) error {
	return s.db.Create(obj).Error
}

func (s service) Save(obj interface{}) error {
	return s.db.Save(obj).Error
}

func (s service) translate(err error) error {
	if err == nil {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return weddinginvitation.RecordNotFound{Model: "record"}
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return weddinginvitation.DuplicateData{Message: "duplicate data"}
	}
	return err
}

func (s service) Find(id uuid.UUID, obj interface{}) error {
	err := s.db.First(obj, id).Error
	return s.translate(err)
}

func (s service) Updates(obj interface{}, updates map[string]interface{}) (int64, error) {
	result := s.db.Model(obj).Updates(updates)
	return result.RowsAffected, result.Error
}

func (s service) Delete(obj interface{}) error {
	err := s.db.Delete(obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &weddinginvitation.RecordNotFound{
			Model: "record",
		}
	}
	return err
}
