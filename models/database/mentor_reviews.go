package database

import (
	"time"

	"github.com/google/uuid"
)

type MentorReviews struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	MentorID       uuid.UUID `gorm:"type:uuid;not null"`
	Description    string    `gorm:"type:text; not null"`
	IsActive       bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy      string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy      string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt      time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime;not null;default:now()"`
}
