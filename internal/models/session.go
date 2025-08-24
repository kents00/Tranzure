package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionID  uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"session_id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	IPAddress  string    `gorm:"type:varchar(45);not null" json:"ip_address"`
	DeviceInfo string    `gorm:"type:text" json:"device_info"`
	LoginAt    time.Time `gorm:"autoCreateTime" json:"login_at"`
	ExpiresAt  time.Time `gorm:"not null" json:"expires_at"`

	// Relationships
	User User `gorm:"foreignKey:UserID;references:UserID" json:"user,omitempty"`
}

func (Session) TableName() string {
	return "sessions"
}
