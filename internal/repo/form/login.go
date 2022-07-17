package form

import (
	"database/sql"

	"github.com/fiber-go-pos-app/utils/pkg/postgres"
	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
	formEntity "github.com/fiber-go-pos-app/internal/entity/form"
)

const queryLoginForm = `
	SELECT user_name, password
	FROM users
	WHERE user_name = $1
	LIMIT 1
`

func LoginFormRepo(ctx *fiber.Ctx, userName string) (formEntity.LoginRequest, error) {
	var user formEntity.LoginRequest
	db := postgres.GetPgConn()

	if err := db.GetContext(ctx.Context(), &user, queryLoginForm, userName); err != nil {
		if err == sql.ErrNoRows {
			return user, constantsEntity.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}
