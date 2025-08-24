package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models"
	"github.com/kento/tranzure/internal/models/validation"
)

func TestKYCValidation(t *testing.T) {
	// Create a valid KYC verification as baseline
	now := time.Now()
	reason := "Missing information"
	validKYC := models.KYCVerification{
		KYCID:        uuid.New(),
		UserID:       uuid.New(),
		Status:       models.KYCStatusPending,
		DocumentType: "passport",
		SubmittedAt:  time.Now(),
	}

	// Test the valid KYC verification
	if err := validKYC.Validate(); err != nil {
		t.Errorf("Expected valid KYC to pass validation, got error: %v", err)
	}

	// Define test cases with invalid data
	testCases := []struct {
		name          string
		modifyKYC     func(*models.KYCVerification)
		expectedError error
	}{
		{
			name: "Empty KYCID",
			modifyKYC: func(k *models.KYCVerification) {
				k.KYCID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty UserID",
			modifyKYC: func(k *models.KYCVerification) {
				k.UserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty Document Type",
			modifyKYC: func(k *models.KYCVerification) {
				k.DocumentType = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid Status",
			modifyKYC: func(k *models.KYCVerification) {
				k.Status = "invalid-status"
			},
			expectedError: validation.ErrInvalidValue,
		},
		{
			name: "Rejected Without Reason",
			modifyKYC: func(k *models.KYCVerification) {
				k.Status = models.KYCStatusRejected
				k.ReviewedAt = &now
				k.RejectionReason = nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Approved Without ReviewedAt",
			modifyKYC: func(k *models.KYCVerification) {
				k.Status = models.KYCStatusApproved
				k.ReviewedAt = nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Rejected Without ReviewedAt",
			modifyKYC: func(k *models.KYCVerification) {
				k.Status = models.KYCStatusRejected
				k.RejectionReason = &reason
				k.ReviewedAt = nil
			},
			expectedError: validation.ErrEmptyField,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the valid KYC verification
			kyc := validKYC

			// Apply the modification to make it invalid
			tc.modifyKYC(&kyc)

			// Validate and check for expected error
			err := kyc.Validate()
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
