package userservice

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetProfile(userID string) (*UserProfile, error) {
	var profile UserProfile
	if err := s.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}
	return &profile, nil
}

func (s *UserService) UpdateProfile(userID string, req UpdateProfileRequest) (*UserProfile, error) {
	// Get existing profile or create new
	var profile UserProfile
	if err := s.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create new profile
			profile = UserProfile{
				ID:              uuid.New().String(),
				UserID:          userID,
				DateOfBirth:     req.DateOfBirth,
				Country:         req.Country,
				City:            req.City,
				Address:         req.Address,
				PostalCode:      req.PostalCode,
				PhoneNumber:     req.PhoneNumber,
				Occupation:      req.Occupation,
				BioDetails:      req.BioDetails,
			}
			if err := s.db.Create(&profile).Error; err != nil {
				return nil, err
			}
			return &profile, nil
		}
		return nil, err
	}

	// Update existing profile
	profile.DateOfBirth = req.DateOfBirth
	profile.Country = req.Country
	profile.City = req.City
	profile.Address = req.Address
	profile.PostalCode = req.PostalCode
	profile.PhoneNumber = req.PhoneNumber
	profile.Occupation = req.Occupation
	profile.BioDetails = req.BioDetails

	if err := s.db.Save(&profile).Error; err != nil {
		return nil, err
	}

	return &profile, nil
}
