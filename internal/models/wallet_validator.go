package models

import (
	"fmt"

	"github.com/kento/tranzure/internal/models/validation"
)

// Validate performs validation on the Wallet model
func (w *Wallet) Validate() error {
	// Check required fields
	if validation.IsEmptyUUID(w.WalletID) {
		return fmt.Errorf("wallet_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(w.UserID) {
		return fmt.Errorf("user_id: %w", validation.ErrEmptyField)
	}

	// Validate wallet type
	if w.Type != WalletTypeCrypto && w.Type != WalletTypeFiat {
		return fmt.Errorf("type: %w", validation.ErrInvalidValue)
	}

	// Validate currency
	if w.Currency == "" {
		return fmt.Errorf("currency: %w", validation.ErrEmptyField)
	}

	if !validation.IsValidCurrency(w.Currency) {
		return fmt.Errorf("currency: %w", validation.ErrInvalidFormat)
	}

	// For crypto wallets, validate address if not empty
	if w.Type == WalletTypeCrypto && w.Address != "" {
		// Add crypto-specific address validation logic here if needed
		if len(w.Address) < 26 || len(w.Address) > 100 {
			return fmt.Errorf("address: %w", validation.ErrInvalidLength)
		}
	}

	// Ensure balance is not negative
	if validation.IsNegativeAmount(w.Balance) {
		return fmt.Errorf("balance: %w", validation.ErrNegativeAmount)
	}

	return nil
}
