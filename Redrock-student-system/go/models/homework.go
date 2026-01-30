package models

import "time"

type Homework struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null;" json:"title"`
	Description string    `gorm:"type:text;" json:"description"`
	Department  string    `gorm:"type:enum('backend','frontend','sre','product','design','android','ios')" json:"department"`
	CreatorID   uint64    `gorm:"not null;" json:"creator_id"`
	Creator     User      `gorm:"foreignKey:CreatorID" json:"creator,omitempty"`
	Deadline    time.Time `gorm:"" json:"deadline"`
	AllowLate   bool      `gorm:"default:false" json:"allow_late"`
	CreatedAt   time.Time `gorm:"" json:"created_at"`
	UpdatedAt   time.Time `gorm:"" json:"updated_at"`
}
