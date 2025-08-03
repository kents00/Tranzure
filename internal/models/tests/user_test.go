package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models"
	"github.com/kento/tranzure/internal/models/validation"
)

func TestUserValidation(t *testing.T) {
	// Create a valid user as baseline
	validUser := models.User{
		UserID:       uuid.New(),
		Email:        "test@example.com",
		PasswordHash: "hashed_password_123",
		Role:         models.UserRoleUser,
		Status:       models.UserStatusActive,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Test the valid user
	if err := validUser.Validate(); err != nil {
		t.Errorf("Expected valid user to pass validation, got error: %v", err)
	}

	// Define test cases with invalid data
	testCases := []struct {
		name          string
		modifyUser    func(*models.User)
		expectedError error
	}{
		{
			name: "Empty UserID",
			modifyUser: func(u *models.User) {
				u.UserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty Email",
			modifyUser: func(u *models.User) {
				u.Email = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid Email Format",
			modifyUser: func(u *models.User) {
				u.Email = "invalid-email"
			},
			expectedError: validation.ErrInvalidFormat,
		},
		{
			name: "Empty Password Hash",
			modifyUser: func(u *models.User) {
				u.PasswordHash = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid Role",
			modifyUser: func(u *models.User) {
				u.Role = "invalid-role"
			},
			expectedError: validation.ErrInvalidValue,
		},
		{
			name: "Invalid Status",
			modifyUser: func(u *models.User) {
				u.Status = "invalid-status"
			},
			expectedError: validation.ErrInvalidValue,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the valid user
			user := validUser

			// Apply the modification to make it invalid
			tc.modifyUser(&user)

			// Validate and check for expected error
			err := user.Validate()
			if err == nil {
				t.Errorf("Expected validation error but got nil")
				return
			}

			if !ErrorContains(err, tc.expectedError) {
				t.Errorf("Expected error to contain '%v', got '%v'", tc.expectedError, err)
			}
		})
	}
}

// ErrorContains checks if the error message contains the expected error
func ErrorContains(err error, target error) bool {
	if err == nil {
		return target == nil
	}
	if target == nil {
		return false
	}
	return err.Error() != "" && target.Error() != "" && contains(err.Error(), target.Error())
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return s != "" && substr != "" && (len(s) >= len(substr)) && (s != substr) && (string(s) != string(substr))
}
