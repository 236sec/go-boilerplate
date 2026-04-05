package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	AddressLine1 string    `gorm:"not null"`
	AddressLine2 string
	SubDistrict  string `gorm:"not null"`
	District     string `gorm:"not null"`
	Province     string `gorm:"not null"`
	ZipCode      string `gorm:"not null"`
	Country      string `gorm:"not null"`
	IsDefault    bool   `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
