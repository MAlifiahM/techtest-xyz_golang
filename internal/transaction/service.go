package transaction

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"xyz_golang/internal/domain"
)

type TransactionService struct {
	consumerRepo    domain.ConsumerRepository
	transactionRepo domain.TransactionRepository
}

func NewTransactionService(consumer domain.ConsumerRepository, transaction domain.TransactionRepository) domain.TransactionService {
	return &TransactionService{
		consumerRepo:    consumer,
		transactionRepo: transaction,
	}
}

func (s *TransactionService) Store(transaction *domain.Transaction) error {
	consumer, err := s.consumerRepo.GetByID(transaction.ConsumerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("consumer not found")
		}

		return err
	}

	tenorFound := false
	for _, limit := range consumer.Limits {
		if limit.Tenor == transaction.Tenor {
			tenorFound = true
		}
	}

	if !tenorFound {
		return errors.New("tenor not found")
	}

	return s.transactionRepo.Store(transaction)
}

func (s *TransactionService) GetByID(id uuid.UUID) (*domain.Transaction, error) {
	return s.transactionRepo.GetByID(id)
}
