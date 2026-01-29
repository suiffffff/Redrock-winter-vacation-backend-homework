package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64         `gorm:"primaryKey;autoIncrement;" json:"id"`
	Username   string         `gorm:"type:varchar(50);unique;not null;" json:"username"`
	Password   string         `gorm:"type:varchar(255);not null;" json:"-"`
	Nickname   string         `gorm:"type:varchar(50);not null;" json:"nickname"`
	Role       string         `gorm:"type:enum('student','admin');default:'student';" json:"role"`
	Department string         `gorm:"type:enum('backend','frontend','sre','product','design','android','ios')" json:"department"`
	Email      string         `gorm:"type:varchar(100);" json:"email"`
	CreatedAt  time.Time      `gorm:"" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"" json:"-"`
}

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

type Submission struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	HomeworkID uint64   `gorm:"not null" json:"homework_id"`
	Homework   Homework `gorm:"foreignKey:HomeworkID" json:"worker,omitempty"`

	StudentID uint64 `gorm:"not null" json:"student_id"`
	Student   User   `gorm:"foreignKey:StudentID" json:"student"`

	Content     string `gorm:"type:text" json:"content"`
	FileUrl     string `gorm:"type:varchar(500)" json:"file_url"`
	IsLate      bool   `gorm:"" json:"is_late"`
	Score       *int   `gorm:"" json:"score"`
	Comment     string `gorm:"type:text" json:"comment"`
	IsExcellent bool   `gorm:"default:false" json:"is_excellent"`

	ReviewerID *uint64 `gorm:"" json:"reviewer_id"`
	Reviewer   User    `gorm:"foreignKey:ReviewerID" json:"reviewer"`

	SubmittedAt time.Time `gorm:"" json:"submitted_at"`
	ReviewedAt  time.Time `gorm:"" json:"reviewed_at"`
	CreatedAt   time.Time `gorm:"" json:"created_at"`
	UpdatedAt   time.Time `gorm:"" json:"updated_at"`
}
