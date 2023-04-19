package models

import (
	"time"
)

type Course struct {
	CourseId    int       `gorm:"not null;uniqueIndex;primaryKey;" json:"course_id"`
	Title       string    `gorm:"not null;size:256;" json:"title"`
	Slug        string    `gorm:"not null;size:256;" json:"slug"`
	Description string    `gorm:"not null;" json:"description"`
	Price       int       `gorm:"not null;" json:"price"`
	Level       string    `gorm:"not null;" json:"level"`                  // Beginner intermediate Advance
	Language    string    `gorm:"not null;" json:"language"`               // Indonesia English
	
	// Course has many UserCourses, CourseId is the foreign key
	UserCourses []UserCourse

	CreatedAt   time.Time `gorm:"not null;->;<-:create" json:"created_at"` // allow read and create
	UpdatedAt   time.Time `gorm:"not null;" json:"updated_at"`
}
