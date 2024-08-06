package model

import (
	"time"
	
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FileName string `gorm:"size:255;not null" json:"file_name"`
	Type string `gorm:"size:100;not null" json:"type"`
	Url string `gorm:"size:255;not null" json:"url"`
	Path string `gorm:"size:255;not null" json:"path"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}