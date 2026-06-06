package authservice

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	Email         string    `gorm:"uniqueIndex" json:"email"`
	Username      string    `gorm:"uniqueIndex" json:"username"`
	PasswordHash  string    `json:"-"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	KYCLevel      int       `json:"kyc_level"`
	KYCStatus     string    `json:"kyc_status"`
	TwoFAEnabled  bool      `json:"two_fa_enabled"`
	TwoFASecret   string    `json:"-"`
	Role          string    `json:"role"`
	Status        string    `json:"status"`
	LastLoginAt   *time.Time `json:"last_login_at"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type RefreshToken struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `gorm:"uniqueIndex" json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	TwoFA    string `json:"two_fa,omitempty"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Username  string `json:"username" binding:"required,min=3"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type RegisterResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
