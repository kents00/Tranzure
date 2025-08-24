package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	LogID      uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"log_id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Action     string    `gorm:"type:varchar(255);not null" json:"action"`
	TargetType string    `gorm:"type:varchar(50);not null" json:"target_type"`
	TargetID   uuid.UUID `gorm:"type:uuid;not null" json:"target_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
