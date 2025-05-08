package domain

import (
	"github.com/google/uuid"
	"time"
)

type Consumer struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	NIK          string    `json:"nik" gorm:"uniqueIndex;not null"`
	FullName     string    `json:"full_name" gorm:"not null"`
	LegalName    string    `json:"legal_name"`
	PlaceOfBirth string    `json:"place_of_birth" gorm:"not null"`
	DateOfBirth  time.Time `json:"date_of_birth" gorm:"not null"`
	Salary       float64   `json:"salary"`
	PhotoKTP     string    `json:"photo_ktp"`
	PhotoSelfie  string    `json:"photo_selfie"`
	Limits       []Limit   `json:"limits" gorm:"foreignKey:ConsumerID"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
