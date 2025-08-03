package models

import (
	"time"
	"github.com/google/uuid"
)

type UserRole string
type UserStatus string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

const (
	UserStatusActive UserStatus = "active"
	UserStatusBanned UserStatus = "banned"
)

type User struct {
	UserID       uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"user_id"`
	Email        string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	PasswordHash string     `gorm:"type:varchar(255);not null" json:"-"`
	Role         UserRole   `gorm:"type:varchar(10);not null;default:'user'" json:"role"`
	Status       UserStatus `gorm:"type:varchar(10);not null;default:'active'" json:"status"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}