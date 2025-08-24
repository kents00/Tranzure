package models

import (
	"time"

	"github.com/google/uuid"
)

type KYCStatus string

const (
	KYCStatusPending  KYCStatus = "pending"
	KYCStatusApproved KYCStatus = "approved"
	KYCStatusRejected KYCStatus = "rejected"
)

type KYCVerification struct {
	KYCID           uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"kyc_id"`
	UserID          uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Status          KYCStatus  `gorm:"type:varchar(10);not null;default:'pending'" json:"status"`
	DocumentType    string     `gorm:"type:varchar(50);not null" json:"document_type"`
	SubmittedAt     time.Time  `gorm:"autoCreateTime" json:"submitted_at"`
	ReviewedAt      *time.Time `gorm:"default:null" json:"reviewed_at,omitempty"`
	RejectionReason *string    `gorm:"type:text" json:"rejection_reason,omitempty"`

	// Relationships
	User User `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`
}

func (KYCVerification) TableName() string {
	return "kyc_verifications"
}
