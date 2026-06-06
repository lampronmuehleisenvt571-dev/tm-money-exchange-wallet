package authservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/lampronmuehleisenvt571-dev/tm-money-exchange-wallet/src/shared/auth"
)

type AuthService struct {
	db         *gorm.DB
	jwtManager *auth.JWTManager
}

func NewAuthService(db *gorm.DB, jwtManager *auth.JWTManager) *AuthService {
	return &AuthService{
		db:         db,
		jwtManager: jwtManager,
	}
}

func (s *AuthService) Register(req RegisterRequest) (*RegisterResponse, error) {
	// Check if user already exists
	var existingUser User
	if err := s.db.Where("email = ? OR username = ?", req.Email, req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: string(hash),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		KYCLevel:     0,
		KYCStatus:    "pending",
		Role:         "user",
		Status:       "active",
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &RegisterResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Message:  "User registered successfully. Please verify your email.",
	}, nil
}

func (s *AuthService) Login(req LoginRequest) (*LoginResponse, error) {
	// Find user by email
	var user User
	if err := s.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check 2FA if enabled
	if user.TwoFAEnabled && req.TwoFA == "" {
		return nil, errors.New("2fa code required")
	}

	// Generate tokens
	accessToken, err := s.jwtManager.GenerateToken(user.ID, user.Email, user.Username, user.KYCLevel, user.Role)
	if err != nil {
		return nil, err
	}

	// Create refresh token
	refreshTokenStr := uuid.New().String()
	refreshToken := RefreshToken{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Token:     refreshTokenStr,
		ExpiresAt: time.Now().AddDate(0, 0, 7), // 7 days
	}

	if err := s.db.Create(&refreshToken).Error; err != nil {
		return nil, err
	}

	// Update last login
	s.db.Model(&user).Update("last_login_at", time.Now())

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.jwtManager.tokenTTL.Seconds()),
	}, nil
}

func (s *AuthService) RefreshToken(refreshTokenStr string) (*LoginResponse, error) {
	// Find refresh token
	var refreshToken RefreshToken
	if err := s.db.Where("token = ?", refreshTokenStr).First(&refreshToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid refresh token")
		}
		return nil, err
	}

	// Check if expired
	if time.Now().After(refreshToken.ExpiresAt) {
		return nil, errors.New("refresh token expired")
	}

	// Find user
	var user User
	if err := s.db.Where("id = ?", refreshToken.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	// Generate new access token
	accessToken, err := s.jwtManager.GenerateToken(user.ID, user.Email, user.Username, user.KYC Level, user.Role)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		TokenType:    "Bearer",
		ExpiresIn:    int(s.jwtManager.tokenTTL.Seconds()),
	}, nil
}

func (s *AuthService) Logout(userID string) error {
	return s.db.Where("user_id = ?", userID).Delete(&RefreshToken{}).Error
}
