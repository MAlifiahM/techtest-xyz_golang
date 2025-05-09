package consumer

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"xyz_golang/internal/domain"
)

type pgsqlUserRepository struct {
	db *gorm.DB
}

func NewPgsqlUserRepository(db *gorm.DB) domain.ConsumerRepository {
	return &pgsqlUserRepository{db: db}
}

func (r *pgsqlUserRepository) GetByID(id uuid.UUID) (*domain.Consumer, error) {
	var consumer domain.Consumer
	if err := r.db.Preload("Limits").First(&consumer, id).Error; err != nil {
		return nil, err
	}
	return &consumer, nil
}

func (r *pgsqlUserRepository) Store(consumer *domain.Consumer) error {
	return r.db.Create(consumer).Error
}
