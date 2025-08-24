package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models"
	"github.com/kento/tranzure/internal/models/validation"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

func TestTransactionValidation(t *testing.T) {
	// Create a valid transaction as baseline
	txHash := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"
	validTransaction := models.Transaction{
		TxID:       uuid.New(),
		FromUserID: uuid.New(),
		ToUserID:   uuid.New(),
		WalletID:   uuid.New(),
		Type:       models.TransactionTypeCrypto,
		Currency:   "BTC",
		Amount:     decimal.NewFromFloat(0.5),
		Status:     models.TransactionStatusPending,
		TxHash:     &txHash,
		Metadata:   datatypes.JSON([]byte(`{"fee": "0.001", "network": "mainnet"}`)),
		CreatedAt:  time.Now(),
	}

	// Test the valid transaction
	if err := validTransaction.Validate(); err != nil {
		t.Errorf("Expected valid transaction to pass validation, got error: %v", err)
	}

	// Define test cases with invalid data
	testCases := []struct {
		name              string
		modifyTransaction func(*models.Transaction)
		expectedError     error
	}{
		{
			name: "Empty TxID",
			modifyTransaction: func(tx *models.Transaction) {
				tx.TxID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty FromUserID",
			modifyTransaction: func(tx *models.Transaction) {
				tx.FromUserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty ToUserID",
			modifyTransaction: func(tx *models.Transaction) {
				tx.ToUserID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Empty WalletID",
			modifyTransaction: func(tx *models.Transaction) {
				tx.WalletID = uuid.Nil
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Invalid Transaction Type",
			modifyTransaction: func(tx *models.Transaction) {
				tx.Type = "invalid-type"
			},
			expectedError: validation.ErrInvalidValue,
		},
		{
			name: "Empty Currency",
			modifyTransaction: func(tx *models.Transaction) {
				tx.Currency = ""
			},
			expectedError: validation.ErrEmptyField,
		},
		{
			name: "Zero Amount",
			modifyTransaction: func(tx *models.Transaction) {
				tx.Amount = decimal.NewFromFloat(0)
			},
			expectedError: validation.ErrZeroAmount,
		},
		{
			name: "Invalid Status",
			modifyTransaction: func(tx *models.Transaction) {
				tx.Status = "invalid-status"
			},
			expectedError: validation.ErrInvalidValue,
		},
		{
			name: "Invalid TX Hash Length",
			modifyTransaction: func(tx *models.Transaction) {
				shortHash := "short"
				tx.TxHash = &shortHash
			},
			expectedError: validation.ErrInvalidLength,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the valid transaction
			tx := validTransaction

			// Apply the modification to make it invalid
			tc.modifyTransaction(&tx)

			// Validate and check for expected error
			err := tx.Validate()
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
