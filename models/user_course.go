package models

import "time"

type UserCourse struct {
	UserId    int       `gorm:"not null;" json:"user_id"`
	CourseId  int       `gorm:"not null;uniqueIndex;" json:"course_id"`
	Score     int       `gorm:"null" json:"score"`
	CreatedAt time.Time `gorm:"not null;->;<-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;" json:"updated_at"`
}
