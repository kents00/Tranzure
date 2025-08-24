package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models"
	"github.com/kento/tranzure/internal/models/validation"
)

func TestSessionValidation(t *testing.T) {
	// Create a valid session as baseline
	validSession := models.Session{
		SessionID:  uuid.New(),
		UserID:     uuid.New(),
		IPAddress:  "192.168.1.1",
		DeviceInfo: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
		LoginAt:    time.Now(),
		ExpiresAt:  time.Now().Add(24 * time.Hour),
	}

	// Test the valid session
	if err := validSession.Validate(); err != nil {
		t.Errorf("Expected valid session to pass validation, got error: %v", err)
	}

	// Define test cases with invalid data
	testCases := []struct {
		name          string
		modifySession func(*models.Session)
		expectedError error
	}{
		{
			name: "Empty SessionID",
			modifySession: func(s *models.Session) {
				s.SessionID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty UserID",
			modifySession: func(s *models.Session) {
				s.UserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty IP Address",
			modifySession: func(s *models.Session) {
				s.IPAddress = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid IP Address",
			modifySession: func(s *models.Session) {
				s.IPAddress = "invalid.ip"
			},
			expectedError: validation.ErrInvalidFormat,
		},
		{
			name: "Zero Expiration Time",
			modifySession: func(s *models.Session) {
				s.ExpiresAt = time.Time{}
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Past Expiration Time",
			modifySession: func(s *models.Session) {
				s.ExpiresAt = time.Now().Add(-24 * time.Hour)
			},
			expectedError: validation.ErrPastDate,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the valid session
			session := validSession

			// Apply the modification to make it invalid
			tc.modifySession(&session)

			// Validate and check for expected error
			err := session.Validate()
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
