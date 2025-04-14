package database

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.NullUUID `gorm:"type:uuid,primary_key;column:id;" json:"id"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt time.Time     `gorm:"index;column:deleted_at" json:"deletedAt"`
}
