package form

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name" validate:"required,max=30"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponse struct {
	UserID   string `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
	JWTToken string `json:"jwt_token" form:"jwt_token"`
}
