package userservice

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	UserID      string    `gorm:"uniqueIndex" json:"user_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	PostalCode  string    `json:"postal_code"`
	PhoneNumber string    `json:"phone_number"`
	Occupation  string    `json:"occupation"`
	BioDetails  string    `json:"bio_details"`
	Verified    bool      `json:"verified"`
	VerifiedAt  *time.Time `json:"verified_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type UpdateProfileRequest struct {
	DateOfBirth time.Time `json:"date_of_birth"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	PostalCode  string    `json:"postal_code"`
	PhoneNumber string    `json:"phone_number"`
	Occupation  string    `json:"occupation"`
	BioDetails  string    `json:"bio_details"`
}

type UserProfileResponse struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Verified    bool      `json:"verified"`
}
