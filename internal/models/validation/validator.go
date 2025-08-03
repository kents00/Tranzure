package validation

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Common validation errors
var (
	ErrEmptyField     = errors.New("required field is empty")
	ErrInvalidFormat  = errors.New("field has invalid format")
	ErrInvalidValue   = errors.New("field has invalid value")
	ErrInvalidLength  = errors.New("field has invalid length")
	ErrNegativeAmount = errors.New("amount cannot be negative")
	ErrZeroAmount     = errors.New("amount cannot be zero")
	ErrPastDate       = errors.New("date is in the past")
	ErrFutureDate     = errors.New("date is in the future")
)

// Validator defines the interface that all model validators should implement
type Validator interface {
	Validate() error
}

// IsEmptyUUID checks if a UUID is nil or zero value
func IsEmptyUUID(id uuid.UUID) bool {
	return id == uuid.Nil
}

// IsValidEmail checks if the provided string is a valid email address
func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// IsValidCurrency checks if the provided currency code is valid
func IsValidCurrency(currency string) bool {
	// Simple check - in production would validate against a list of currency codes
	return len(currency) >= 2 && len(currency) <= 10
}

// IsNegativeAmount checks if the amount is negative
func IsNegativeAmount(amount decimal.Decimal) bool {
	return amount.IsNegative()
}

// IsZeroAmount checks if the amount is zero
func IsZeroAmount(amount decimal.Decimal) bool {
	return amount.IsZero()
}

// IsValidIPAddress checks if the provided string is a valid IP address
func IsValidIPAddress(ip string) bool {
	// Simple IP validation - would use net.ParseIP in production
	ipv4Pattern := `^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`
	ipv6Pattern := `^(([0-9a-fA-F]{1,4}:){7}([0-9a-fA-F]{1,4}|:))|(([0-9a-fA-F]{1,4}:){6}(:[0-9a-fA-F]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9a-fA-F]{1,4}:){5}(((:[0-9a-fA-F]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9a-fA-F]{1,4}:){4}(((:[0-9a-fA-F]{1,4}){1,3})|((:[0-9a-fA-F]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9a-fA-F]{1,4}:){3}(((:[0-9a-fA-F]{1,4}){1,4})|((:[0-9a-fA-F]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9a-fA-F]{1,4}:){2}(((:[0-9a-fA-F]{1,4}){1,5})|((:[0-9a-fA-F]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9a-fA-F]{1,4}:){1}(((:[0-9a-fA-F]{1,4}){1,6})|((:[0-9a-fA-F]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9a-fA-F]{1,4}){1,7})|((:[0-9a-fA-F]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))$`

	reIPv4 := regexp.MustCompile(ipv4Pattern)
	reIPv6 := regexp.MustCompile(ipv6Pattern)

	return reIPv4.MatchString(ip) || reIPv6.MatchString(ip)
}

// IsPastTime checks if the provided time is in the past
func IsPastTime(t time.Time) bool {
	return t.Before(time.Now())
}

// IsFutureTime checks if the provided time is in the future
func IsFutureTime(t time.Time) bool {
	return t.After(time.Now())
}
