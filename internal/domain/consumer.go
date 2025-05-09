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

type ConsumerRequest struct {
	NIK          string  `json:"nik" validate:"required"`
	FullName     string  `json:"full_name" validate:"required"`
	LegalName    string  `json:"legal_name"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required" faker:"date"`
	Salary       float64 `json:"salary"`
	PhotoKTP     string  `json:"photo_ktp"`
	PhotoSelfie  string  `json:"photo_selfie"`
}

type ConsumerRepository interface {
	GetByID(id uuid.UUID) (*Consumer, error)
	Store(consumer *Consumer) error
}

type ConsumerService interface {
	GetByID(id uuid.UUID) (*Consumer, error)
	GetLimit(id uuid.UUID) (*[]Limit, error)
	Store(consumer *Consumer) error
	StoreLimit(limit *Limit) error
}
