package params

type EnrollCourse struct {
	UserId   int `json:"user_id" validate:"required"`
	CourseId int `json:"course_id" validate:"required"`
}
