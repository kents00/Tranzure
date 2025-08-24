package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models"
	"github.com/kento/tranzure/internal/models/validation"
)

func TestAuditLogValidation(t *testing.T) {
	// Create a valid audit log as baseline
	validAuditLog := models.AuditLog{
		LogID:      uuid.New(),
		UserID:     uuid.New(),
		Action:     "user.login",
		TargetType: "session",
		TargetID:   uuid.New(),
		CreatedAt:  time.Now(),
	}

	// Test the valid audit log
	if err := validAuditLog.Validate(); err != nil {
		t.Errorf("Expected valid audit log to pass validation, got error: %v", err)
	}

	// Define test cases with invalid data
	testCases := []struct {
		name           string
		modifyAuditLog func(*models.AuditLog)
		expectedError  error
	}{
		{
			name: "Empty LogID",
			modifyAuditLog: func(a *models.AuditLog) {
				a.LogID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty UserID",
			modifyAuditLog: func(a *models.AuditLog) {
				a.UserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty Action",
			modifyAuditLog: func(a *models.AuditLog) {
				a.Action = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty TargetType",
			modifyAuditLog: func(a *models.AuditLog) {
				a.TargetType = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty TargetID",
			modifyAuditLog: func(a *models.AuditLog) {
				a.TargetID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the valid audit log
			auditLog := validAuditLog

			// Apply the modification to make it invalid
			tc.modifyAuditLog(&auditLog)

			// Validate and check for expected error
			err := auditLog.Validate()
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
