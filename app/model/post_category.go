package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostCategory struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	PostID uuid.UUID `json:"-"`
	CategoryID uuid.UUID `json:"-"`
}