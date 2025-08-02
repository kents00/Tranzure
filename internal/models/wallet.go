package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type WalletType string

const (
	WalletTypeCrypto WalletType = "crypto"
	WalletTypeFiat   WalletType = "fiat"
)

type Wallet struct {
	WalletID  uuid.UUID       `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"wallet_id"`
	UserID    uuid.UUID       `gorm:"type:uuid;not null;index" json:"user_id"`
	Type      WalletType      `gorm:"type:varchar(10);not null" json:"type"`
	Currency  string          `gorm:"type:varchar(10);not null" json:"currency"`
	Address   string          `gorm:"type:varchar(255)" json:"address"`
	Balance   decimal.Decimal `gorm:"type:decimal(20,8);default:0" json:"balance"`
	IsPrimary bool            `gorm:"default:false" json:"is_primary"`

	// Relationships
	User User `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`
}

func (Wallet) TableName() string {
	return "wallets"
}
