package models

import (
	"fmt"

	"github.com/kento/tranzure/internal/models/validation"
)

// Validate performs validation on the KYC model
func (k *KYCVerification) Validate() error {
	// Check required fields
	if validation.IsEmptyUUID(k.KYCID) {
		return fmt.Errorf("kyc_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(k.UserID) {
		return fmt.Errorf("user_id: %w", validation.ErrEmptyField)
	}

	// Validate document type
	if k.DocumentType == "" {
		return fmt.Errorf("document_type: %w", validation.ErrEmptyField)
	}

	// Validate KYC status
	if k.Status != KYCStatusPending &&
		k.Status != KYCStatusApproved &&
		k.Status != KYCStatusRejected {
		return fmt.Errorf("status: %w", validation.ErrInvalidValue)
	}

	// If status is rejected, ensure rejection reason is provided
	if k.Status == KYCStatusRejected && (k.RejectionReason == nil || *k.RejectionReason == "") {
		return fmt.Errorf("rejection_reason: %w", validation.ErrEmptyField)
	}

	// If status is approved or rejected, ensure reviewed_at is provided
	if (k.Status == KYCStatusApproved || k.Status == KYCStatusRejected) && k.ReviewedAt == nil {
		return fmt.Errorf("reviewed_at: %w", validation.ErrEmptyField)
	}

	return nil
}
