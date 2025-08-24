package models

import (
	"fmt"

	"github.com/kento/tranzure/internal/models/validation"
)

// Validate performs validation on the Session model
func (s *Session) Validate() error {
	// Check required fields
	if validation.IsEmptyUUID(s.SessionID) {
		return fmt.Errorf("session_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(s.UserID) {
		return fmt.Errorf("user_id: %w", validation.ErrEmptyField)
	}

	// Validate IP address
	if s.IPAddress == "" {
		return fmt.Errorf("ip_address: %w", validation.ErrEmptyField)
	}

	if !validation.IsValidIPAddress(s.IPAddress) {
		return fmt.Errorf("ip_address: %w", validation.ErrInvalidFormat)
	}

	// Validate expiration time
	if s.ExpiresAt.IsZero() {
		return fmt.Errorf("expires_at: %w", validation.ErrEmptyField)
	}

	if validation.IsPastTime(s.ExpiresAt) {
		return fmt.Errorf("expires_at: %w", validation.ErrPastDate)
	}

	return nil
}
