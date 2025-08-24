package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kento/tranzure/internal/models/validation"
	"github.com/shopspring/decimal"
)

func TestIsValidEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag@example.co.uk", true},
		{"", false},
		{"invalid", false},
		{"invalid@", false},
		{"@invalid.com", false},
		{"test@.com", false},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			result := validation.IsValidEmail(tc.email)
			if result != tc.expected {
				t.Errorf("IsValidEmail(%q) = %v; want %v", tc.email, result, tc.expected)
			}
		})
	}
}

func TestIsEmptyUUID(t *testing.T) {
	testCases := []struct {
		name     string
		uuid     uuid.UUID
		expected bool
	}{
		{"Nil UUID", uuid.Nil, true},
		{"Valid UUID", uuid.New(), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validation.IsEmptyUUID(tc.uuid)
			if result != tc.expected {
				t.Errorf("IsEmptyUUID(%v) = %v; want %v", tc.uuid, result, tc.expected)
			}
		})
	}
}

func TestIsValidCurrency(t *testing.T) {
	testCases := []struct {
		currency string
		expected bool
	}{
		{"USD", true},
		{"EUR", true},
		{"BTC", true},
		{"ETH", true},
		{"", false},
		{"A", false},
		{"TOOOOOOOLONG", false},
	}

	for _, tc := range testCases {
		t.Run(tc.currency, func(t *testing.T) {
			result := validation.IsValidCurrency(tc.currency)
			if result != tc.expected {
				t.Errorf("IsValidCurrency(%q) = %v; want %v", tc.currency, result, tc.expected)
			}
		})
	}
}

func TestIsNegativeAmount(t *testing.T) {
	testCases := []struct {
		name     string
		amount   decimal.Decimal
		expected bool
	}{
		{"Positive", decimal.NewFromFloat(10.0), false},
		{"Zero", decimal.NewFromFloat(0), false},
		{"Negative", decimal.NewFromFloat(-10.0), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validation.IsNegativeAmount(tc.amount)
			if result != tc.expected {
				t.Errorf("IsNegativeAmount(%v) = %v; want %v", tc.amount, result, tc.expected)
			}
		})
	}
}

func TestIsZeroAmount(t *testing.T) {
	testCases := []struct {
		name     string
		amount   decimal.Decimal
		expected bool
	}{
		{"Positive", decimal.NewFromFloat(10.0), false},
		{"Zero", decimal.NewFromFloat(0), true},
		{"Negative", decimal.NewFromFloat(-10.0), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validation.IsZeroAmount(tc.amount)
			if result != tc.expected {
				t.Errorf("IsZeroAmount(%v) = %v; want %v", tc.amount, result, tc.expected)
			}
		})
	}
}

func TestIsValidIPAddress(t *testing.T) {
	testCases := []struct {
		name     string
		ip       string
		expected bool
	}{
		{"IPv4 Valid", "192.168.1.1", true},
		{"IPv4 Invalid", "192.168.1", false},
		{"IPv6 Valid", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", true},
		{"Empty", "", false},
		{"Invalid Format", "not-an-ip", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validation.IsValidIPAddress(tc.ip)
			if result != tc.expected {
				t.Errorf("IsValidIPAddress(%q) = %v; want %v", tc.ip, result, tc.expected)
			}
		})
	}
}

func TestIsPastTime(t *testing.T) {
	testCases := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{"Past", time.Now().Add(-1 * time.Hour), true},
		{"Future", time.Now().Add(1 * time.Hour), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validation.IsPastTime(tc.time)
			if result != tc.expected {
				t.Errorf("IsPastTime(%v) = %v; want %v", tc.time, result, tc.expected)
			}
		})
	}
}

func TestIsFutureTime(t *testing.T) {
	testCases := []struct {
		name     string
		time     time.Time
		expected bool
	}{
		{"Past", time.Now().Add(-1 * time.Hour), false},
		{"Future", time.Now().Add(1 * time.Hour), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := validation.IsFutureTime(tc.time)
			if result != tc.expected {
				t.Errorf("IsFutureTime(%v) = %v; want %v", tc.time, result, tc.expected)
			}
		})
	}
}
