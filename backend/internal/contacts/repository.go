package contacts

import (
	"context"
	"errors"

	"gorm.io/gorm"
)


type ContactRepository interface {
	GetAll(ctx context.Context) ([]Contact, error)
	GetById(ctx context.Context, id uint) (*Contact, error)
	Create(ctx context.Context, contact *Contact) error
	Update(ctx context.Context, id uint, contact *Contact, updates *Contact) error
	Delete(ctx context.Context, id uint) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (cr *contactRepository) GetAll(ctx context.Context) ([]Contact, error) {
	var contacts []Contact

	result := cr.db.WithContext(ctx).Find(&contacts)
	if result.Error != nil {
		return nil, result.Error
	}

	return contact, nil
}

func (cr *contactRepository) GetByID(ctx context.Context, id uint) (*Contact, error) {
	var contact Contact

	result := cr.db.WithContext(ctx).First(&contact, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		
		return nil, result.Error
	}

	return &contact, nil
}

func (cr *contactRepository) Create(ctx context.Context, contact *Contact) error {
	return cr.db.Create(contact).Error
}

func (cr *contactRepository) Update(ctx context.Context, id uint, contact *Contact, updates *Contact) error {
	result := cr.db.WithContext(ctx).Model(contact).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (cr *contactRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.Delete(&Contact{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

