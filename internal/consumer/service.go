package consumer

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"xyz_golang/internal/domain"
)

type ConsumerService struct {
	consumerRepo domain.ConsumerRepository
	limitRepo    domain.LimitRepository
}

func NewConsumerService(consumer domain.ConsumerRepository, limit domain.LimitRepository) domain.ConsumerService {
	return &ConsumerService{
		consumerRepo: consumer,
		limitRepo:    limit,
	}
}

func (s *ConsumerService) GetByID(id uuid.UUID) (*domain.Consumer, error) {
	consumer, err := s.consumerRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
		return nil, err
	}

	return consumer, nil
}

func (s *ConsumerService) GetLimit(id uuid.UUID) (*[]domain.Limit, error) {
	limit, err := s.limitRepo.LimitByConsumerID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fiber.ErrNotFound
		}
	}

	return limit, nil
}

func (s *ConsumerService) Store(consumer *domain.Consumer) error {
	return s.consumerRepo.Store(consumer)
}
