package models

import (
	"time"
)

type User struct {
	UserId    int       `gorm:"not null;uniqueIndex;primaryKey;" json:"user_id"`
	Name      string    `gorm:"not null;size:256" json:"name"`
	Email     string    `gorm:"not null;uniqueIndex;size:256;" json:"email"`
	Password  string    `gorm:"not null;" json:"password"`
	Role      string    `gorm:"not null;size:5" json:"role"`

	// User has many UserCourses, UserId is the foreign key
	UserCourses []UserCourse

	CreatedAt time.Time `gorm:"not null;->;<-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;" json:"updated_at"`
}
