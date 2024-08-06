package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FullName string `gorm:"size:255;not null" json:"full_name"`
	Email string `gorm:"size:100;not null;unique" json:"email"`
	Role string `gorm:"size:100;not null;type:enum('Super Admin', 'Creator')" json:"role"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Status string `gorm:"size:100;not null;type:enum('Active', 'Suspend')" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	AvatarID uuid.UUID `json:"-"`
	Avatar File `gorm:"foreignKey:AvatarID" json:"avatar"`
}