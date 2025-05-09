package limit

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"xyz_golang/internal/domain"
)

type pgsqlLimitRepository struct {
	db *gorm.DB
}

func NewPgsqlLimitRepository(db *gorm.DB) domain.LimitRepository {
	return &pgsqlLimitRepository{db: db}
}

func (r *pgsqlLimitRepository) LimitByConsumerID(id uuid.UUID) (*[]domain.Limit, error) {
	var limit []domain.Limit
	if err := r.db.Find(&limit, "consumer_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &limit, nil
}

func (r *pgsqlLimitRepository) Store(limit *domain.Limit) error {
	return r.db.Create(limit).Error
}
