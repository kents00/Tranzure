package models

import (
	"fmt"

	"github.com/kento/tranzure/internal/models/validation"
)

// Validate performs validation on the AuditLog model
func (a *AuditLog) Validate() error {
	// Check required fields
	if validation.IsEmptyUUID(a.LogID) {
		return fmt.Errorf("log_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(a.UserID) {
		return fmt.Errorf("user_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(a.TargetID) {
		return fmt.Errorf("target_id: %w", validation.ErrEmptyField)
	}

	// Validate action
	if a.Action == "" {
		return fmt.Errorf("action: %w", validation.ErrEmptyField)
	}

	// Validate target type
	if a.TargetType == "" {
		return fmt.Errorf("target_type: %w", validation.ErrEmptyField)
	}

	return nil
}
