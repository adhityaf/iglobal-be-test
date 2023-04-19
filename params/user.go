package params

type CreateUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUser struct {
	UserId int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Role   string `json:"role"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
