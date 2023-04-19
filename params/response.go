package params

import "github.com/adhityaf/iglobal-be-test/models"

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

type Statistic struct {
	TotalUser       int64 `json:"total_user"`
	TotalCourse     int64 `json:"total_course"`
	TotalFreeCourse int64 `json:"total_free_course"`
}

type CourseResponse struct {
	Course *models.Course `json:"course"`
}
