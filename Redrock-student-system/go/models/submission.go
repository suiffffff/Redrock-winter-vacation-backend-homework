package models

import "time"

type Submission struct {
	ID uint64 `gorm:"primaryKey;autoIncrement" json:"id"`

	HomeworkID uint64   `gorm:"not null" json:"homework_id"`
	Homework   Homework `gorm:"foreignKey:HomeworkID" json:"homework,omitempty"`

	StudentID uint64 `gorm:"not null" json:"student_id"`
	Student   User   `gorm:"foreignKey:StudentID" json:"student"`

	Content     string `gorm:"type:text" json:"content"`
	FileUrl     string `gorm:"type:varchar(500)" json:"file_url"`
	IsLate      bool   `gorm:"" json:"is_late"`
	Score       *int   `gorm:"" json:"score"`
	Comment     string `gorm:"type:text" json:"comment"`
	IsExcellent bool   `gorm:"default:false" json:"is_excellent"`

	ReviewerID *uint64 `gorm:"" json:"reviewer_id"`
	Reviewer   *User   `gorm:"foreignKey:ReviewerID" json:"reviewer"`

	SubmittedAt time.Time  `gorm:"" json:"submitted_at"`
	ReviewedAt  *time.Time `gorm:"" json:"reviewed_at"`
	CreatedAt   time.Time  `gorm:"" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"" json:"updated_at"`
}
