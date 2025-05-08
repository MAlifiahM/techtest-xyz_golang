package domain

import (
	"github.com/google/uuid"
	"time"
)

type Limit struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ConsumerID uuid.UUID `json:"consumer_id" gorm:"not null;type:uuid"`
	Tenor      int       `json:"tenor" gorm:"not null"`
	Amount     float64   `json:"amount" gorm:"not null"`
	Consumer   Consumer  `json:"consumer" gorm:"foreignKey:ConsumerID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
