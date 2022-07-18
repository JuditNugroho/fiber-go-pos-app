package form

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name" validate:"required,max=30"`
	Password string `json:"password" form:"password" validate:"required"`
}
