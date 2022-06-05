package personalia

type User struct {
	UserID   string `json:"user_id" db:"user_id" validate:"required,max=20"`
	UserName string `json:"user_name" db:"user_name" validate:"max=30"`
	FullName string `json:"full_name" db:"full_name" validate:"max=255"`
	Password string `json:"password,omitempty" db:"password" validate:"max=255"`
	IsAdmin  bool   `json:"is_admin,omitempty" db:"is_admin"`
}
