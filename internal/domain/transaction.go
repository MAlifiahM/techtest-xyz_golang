package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID             uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ConsumerID     uuid.UUID `json:"consumer_id" gorm:"not null;type:uuid"`
	ContractNumber string    `json:"contract_number" gorm:"not null"`
	Tenor          int       `json:"tenor" gorm:"not null"`
	OTR            float64   `json:"otr" gorm:"not null"`
	AdminFee       float64   `json:"admin_fee"`
	Installment    float64   `json:"installment"`
	Interest       float64   `json:"interest"`
	AssetName      string    `json:"asset_name" gorm:"not null"`
	Consumer       Consumer  `json:"consumer" gorm:"foreignKey:ConsumerID;constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
