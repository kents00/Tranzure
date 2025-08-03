package models

import (
	"fmt"

	"github.com/kento/tranzure/internal/models/validation"
)

// Validate performs validation on the User model
func (u *User) Validate() error {
	// Check required fields
	if validation.IsEmptyUUID(u.UserID) {
		return fmt.Errorf("user_id: %w", validation.ErrEmptyField)
	}

	if u.Email == "" {
		return fmt.Errorf("email: %w", validation.ErrEmptyField)
	}

	if !validation.IsValidEmail(u.Email) {
		return fmt.Errorf("email: %w", validation.ErrInvalidFormat)
	}

	if u.PasswordHash == "" {
		return fmt.Errorf("password_hash: %w", validation.ErrEmptyField)
	}

	// Validate role
	if u.Role != UserRoleUser && u.Role != UserRoleAdmin {
		return fmt.Errorf("role: %w", validation.ErrInvalidValue)
	}

	// Validate status
	if u.Status != UserStatusActive && u.Status != UserStatusBanned {
		return fmt.Errorf("status: %w", validation.ErrInvalidValue)
	}

	return nil
}
