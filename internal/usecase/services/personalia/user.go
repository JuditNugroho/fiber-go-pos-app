package personalia

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	personaliaEntity "github.com/fiber-go-pos-app/internal/entity/personalia"
	customPkg "github.com/fiber-go-pos-app/utils/pkg/custom"

	personaliaRepo "github.com/fiber-go-pos-app/internal/repo/personalia"
)

func GetAllUser(ctx *fiber.Ctx) ([]personaliaEntity.User, error) {
	users, err := personaliaRepo.GetAllUser(ctx)
	if err != nil {
		return []personaliaEntity.User{}, err
	}

	return users, nil
}

func GetUserByUserID(ctx *fiber.Ctx, userID string) (personaliaEntity.User, error) {
	user, found, err := personaliaRepo.GetUserByUserID(ctx, userID)
	if err != nil {
		return personaliaEntity.User{}, err
	}

	if !found {
		return personaliaEntity.User{}, fmt.Errorf("user dengan nama : %s tidak ditemukan", user.UserName)
	}

	return user, nil
}

func InsertUser(ctx *fiber.Ctx, user personaliaEntity.User) error {
	if user.Password != "" {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if err := personaliaRepo.InsertUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx *fiber.Ctx, user personaliaEntity.User) error {
	userDB, err := GetUserByUserID(ctx, user.UserID)
	if err != nil {
		return err
	}

	// replace to existing data
	if user.Password == "" {
		user.Password = userDB.Password
	} else {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if err := personaliaRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx *fiber.Ctx, userID string) error {
	if _, err := GetUserByUserID(ctx, userID); err != nil {
		return err
	}

	if err := personaliaRepo.DeleteUser(ctx, userID); err != nil {
		return err
	}

	return nil
}

func UpsertUser(ctx *fiber.Ctx, user personaliaEntity.User) error {
	userDB, found, err := personaliaRepo.GetUserByUserID(ctx, user.UserID)
	if err != nil {
		return err
	}

	if userDB.Password != user.Password {
		user.Password, _ = customPkg.HashPassword(user.Password)
	}

	if !found {
		return personaliaRepo.InsertUser(ctx, user)
	} else {
		return personaliaRepo.UpdateUser(ctx, user)
	}
}
