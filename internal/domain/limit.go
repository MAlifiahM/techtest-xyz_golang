package domain

import (
	"github.com/google/uuid"
	"time"
)

type Limit struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ConsumerID uuid.UUID `json:"-" gorm:"not null;type:uuid"`
	Tenor      int       `json:"tenor" gorm:"not null"`
	Amount     float64   `json:"amount" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type LimitRequest struct {
	Tenor  int     `json:"tenor" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type LimitRepository interface {
	LimitByConsumerID(id uuid.UUID) (*[]Limit, error)
	Store(limit *Limit) error
}
