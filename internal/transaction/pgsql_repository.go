package transaction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"xyz_golang/internal/domain"
)

type pgsqlTransactionRepository struct {
	db *gorm.DB
}

func NewPgsqlTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &pgsqlTransactionRepository{
		db: db,
	}
}

func (r *pgsqlTransactionRepository) Store(transaction *domain.Transaction) error {
	return r.db.Preload("Consumer").Preload("Consumer.Limits").Create(transaction).Error
}

func (r *pgsqlTransactionRepository) GetByID(id uuid.UUID) (*domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.db.Preload("Consumer").Preload("Consumer.Limits").First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
