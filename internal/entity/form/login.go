package form

type LoginRequest struct {
	UserName string `json:"user_name" db:"user_name" validate:"required,max=30"`
	Password string `json:"password" db:"password" validate:"required"`
}
