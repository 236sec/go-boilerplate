package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	PhoneNumber string
	Role        string `gorm:"type:user_role;default:'CUSTOMER'"`
	IsActive    bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Addresses []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
