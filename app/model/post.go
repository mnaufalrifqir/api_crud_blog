package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title string `gorm:"size:255;not null" json:"title"`
	Description string `gorm:"size:255;not null" json:"description"`
	Status string `gorm:"size:100;not null;type:enum('Draft', 'Published')" json:"status"`
	Slug string `gorm:"size:255;not null" json:"slug"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ThumbnailID uuid.UUID `json:"-"`
	Thumbnail File `gorm:"foreignKey:ThumbnailID" json:"thumbnail"`
	PostCategory []PostCategory `gorm:"foreignKey:PostID" json:"post_category"`
}