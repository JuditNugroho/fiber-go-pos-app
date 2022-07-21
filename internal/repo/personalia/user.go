package personalia

import (
	"database/sql"
	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
	"github.com/fiber-go-pos-app/utils/pkg/databases/postgres"

	"github.com/gofiber/fiber/v2"

	userEntity "github.com/fiber-go-pos-app/internal/entity/personalia"
)

const queryGetAllUser = `
	SELECT user_id, user_name, full_name, password, is_admin
	FROM users
	ORDER BY user_id
`

func GetAllUser(ctx *fiber.Ctx) ([]userEntity.User, error) {
	var users []userEntity.User

	db := postgres.GetPgConn()
	if err := db.SelectContext(ctx.Context(), &users, queryGetAllUser); err != nil {
		return users, err
	}
	return users, nil
}

const queryGetUserByUserID = `
	SELECT user_id, user_name, full_name, password, is_admin
	FROM users
	WHERE user_id = $1
`

func GetUserByUserID(ctx *fiber.Ctx, userID string) (userEntity.User, bool, error) {
	var user userEntity.User

	db := postgres.GetPgConn()

	if err := db.GetContext(ctx.Context(), &user, queryGetUserByUserID, userID); err != nil {
		if err == sql.ErrNoRows {
			return user, false, nil
		}
		return user, false, err
	}
	return user, true, nil
}

const queryGetUserByUserName = `
	SELECT user_id, user_name, full_name, password, is_admin
	FROM users
	WHERE user_name = $1
`

func GetUserByUserName(ctx *fiber.Ctx, userName string) (userEntity.User, error) {
	var user userEntity.User
	db := postgres.GetPgConn()

	if err := db.GetContext(ctx.Context(), &user, queryGetUserByUserName, userName); err != nil {
		if err == sql.ErrNoRows {
			return user, constantsEntity.ErrUserNotFound
		}
		return user, err
	}
	return user, nil
}

const insertUser = `
	INSERT INTO users (user_id, user_name, full_name, password, is_admin)
	VALUES (:user_id, :user_name, :full_name, :password, :is_admin)
`

func InsertUser(ctx *fiber.Ctx, user userEntity.User) error {

	db := postgres.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), insertUser, user)
	if err != nil {
		return err
	}
	return nil
}

const updateUser = `
	UPDATE users SET
		user_name = :user_name,
		full_name = :full_name,
		password = :password,
		is_admin = :is_admin,
		update_time = NOW()
	WHERE user_id = :user_id
`

func UpdateUser(ctx *fiber.Ctx, user userEntity.User) error {

	db := postgres.GetPgConn()
	_, err := db.NamedQueryContext(ctx.Context(), updateUser, user)
	if err != nil {
		return err
	}
	return nil
}

const deleteUser = `
	DELETE FROM users
	WHERE user_id = $1
`

func DeleteUser(ctx *fiber.Ctx, userID string) error {

	db := postgres.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteUser, userID)
	if err != nil {
		return err
	}
	return nil
}
