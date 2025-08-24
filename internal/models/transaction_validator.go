package models

import (
	"fmt"

	"github.com/kento/tranzure/internal/models/validation"
)

// Validate performs validation on the Transaction model
func (t *Transaction) Validate() error {
	// Check required fields
	if validation.IsEmptyUUID(t.TxID) {
		return fmt.Errorf("tx_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(t.FromUserID) {
		return fmt.Errorf("from_user_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(t.ToUserID) {
		return fmt.Errorf("to_user_id: %w", validation.ErrEmptyField)
	}

	if validation.IsEmptyUUID(t.WalletID) {
		return fmt.Errorf("wallet_id: %w", validation.ErrEmptyField)
	}

	// Validate transaction type
	if t.Type != TransactionTypeFiat && t.Type != TransactionTypeCrypto {
		return fmt.Errorf("type: %w", validation.ErrInvalidValue)
	}

	// Validate currency
	if t.Currency == "" {
		return fmt.Errorf("currency: %w", validation.ErrEmptyField)
	}

	if !validation.IsValidCurrency(t.Currency) {
		return fmt.Errorf("currency: %w", validation.ErrInvalidFormat)
	}

	// Validate amount
	if validation.IsZeroAmount(t.Amount) {
		return fmt.Errorf("amount: %w", validation.ErrZeroAmount)
	}

	// Validate transaction status
	if t.Status != TransactionStatusPending &&
		t.Status != TransactionStatusSuccess &&
		t.Status != TransactionStatusFailed {
		return fmt.Errorf("status: %w", validation.ErrInvalidValue)
	}

	// For crypto transactions, validate tx_hash if not empty
	if t.Type == TransactionTypeCrypto && t.TxHash != nil && *t.TxHash != "" {
		// Add crypto-specific tx hash validation logic here if needed
		if len(*t.TxHash) < 10 || len(*t.TxHash) > 100 {
			return fmt.Errorf("tx_hash: %w", validation.ErrInvalidLength)
		}
	}

	return nil
}
