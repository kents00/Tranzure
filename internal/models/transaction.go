package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
)

type (
	TransactionType   string
	TransactionStatus string
)

const (
	TransactionTypeFiat   TransactionType = "fiat"
	TransactionTypeCrypto TransactionType = "crypto"
)

const (
	TransactionStatusPending TransactionStatus = "pending"
	TransactionStatusSuccess TransactionStatus = "success"
	TransactionStatusFailed  TransactionStatus = "failed"
)

type Transaction struct {
	TxID       uuid.UUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"tx_id"`
	FromUserID uuid.UUID         `gorm:"type:uuid;not null;index" json:"from_user_id"`
	ToUserID   uuid.UUID         `gorm:"type:uuid;not null;index" json:"to_user_id"`
	WalletID   uuid.UUID         `gorm:"type:uuid;not null;index" json:"wallet_id"`
	Type       TransactionType   `gorm:"type:varchar(10);not null" json:"type"`
	Currency   string            `gorm:"type:varchar(10);not null" json:"currency"`
	Amount     decimal.Decimal   `gorm:"type:decimal(20,8);not null" json:"amount"`
	Status     TransactionStatus `gorm:"type:varchar(10);not null;default:'pending'" json:"status"`
	TxHash     *string           `gorm:"type:varchar(255)" json:"tx_hash,omitempty"`
	Metadata   datatypes.JSON    `gorm:"type:json" json:"metadata,omitempty"`
	CreatedAt  time.Time         `gorm:"autoCreateTime" json:"created_at"`

	// Relationships
	FromUser User   `gorm:"foreignKey:FromUserID;references:UserID" json:"from_user,omitempty"`
	ToUser   User   `gorm:"foreignKey:ToUserID;references:UserID" json:"to_user,omitempty"`
	Wallet   Wallet `gorm:"foreignKey:WalletID;references:WalletID" json:"wallet,omitempty"`
}

func (Transaction) TableName() string {
	return "transactions"
}
