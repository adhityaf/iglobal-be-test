package params

type CreateCourse struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Level       string `json:"level" validate:"required"`
	Language    string `json:"language" validate:"required"`
}

type UpdateCourse struct {
	CourseId    int    `json:"course_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Level       string `json:"level" validate:"required"`
	Language    string `json:"language" validate:"required"`
}

type CourseSearchCriteria struct {
	Query    string `json:"query"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Level    string `json:"level"`
	Language string `json:"language"`
}

func (c *CourseSearchCriteria) SetDefaultValue() {
	if c.Page == 0 {
		c.Page = 1
	}
	if c.Size == 0 {
		c.Size = 10
	}
	if c.Level == ""{
		c.Level = "Beginner"
	}
	if c.Language == ""{
		c.Language = "Indonesia"
	}
}