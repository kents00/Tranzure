package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models"
	"github.com/kento/tranzure/internal/models/validation"
	"github.com/shopspring/decimal"
)

func TestWalletValidation(t *testing.T) {
	// Create a valid wallet as baseline
	validWallet := models.Wallet{
		WalletID:  uuid.New(),
		UserID:    uuid.New(),
		Type:      models.WalletTypeCrypto,
		Currency:  "BTC",
		Address:   "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh",
		Balance:   decimal.NewFromFloat(1.5),
		IsPrimary: true,
	}

	// Test the valid wallet
	if err := validWallet.Validate(); err != nil {
		t.Errorf("Expected valid wallet to pass validation, got error: %v", err)
	}

	// Define test cases with invalid data
	testCases := []struct {
		name          string
		modifyWallet  func(*models.Wallet)
		expectedError error
	}{
		{
			name: "Empty WalletID",
			modifyWallet: func(w *models.Wallet) {
				w.WalletID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty UserID",
			modifyWallet: func(w *models.Wallet) {
				w.UserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid Wallet Type",
			modifyWallet: func(w *models.Wallet) {
				w.Type = "invalid-type"
			},
			expectedError: validation.ErrInvalidValue,
		},
		{
			name: "Empty Currency",
			modifyWallet: func(w *models.Wallet) {
				w.Currency = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid Currency Format",
			modifyWallet: func(w *models.Wallet) {
				w.Currency = "TOOOOOOOLONG"
			},
			expectedError: validation.ErrInvalidFormat,
		},
		{
			name: "Negative Balance",
			modifyWallet: func(w *models.Wallet) {
				w.Balance = decimal.NewFromFloat(-10.0)
			},
			expectedError: validation.ErrNegativeAmount,
		},
		{
			name: "Invalid Crypto Address Length",
			modifyWallet: func(w *models.Wallet) {
				w.Type = models.WalletTypeCrypto
				w.Address = "too-short"
			},
			expectedError: validation.ErrInvalidLength,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the valid wallet
			wallet := validWallet

			// Apply the modification to make it invalid
			tc.modifyWallet(&wallet)

			// Validate and check for expected error
			err := wallet.Validate()
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
