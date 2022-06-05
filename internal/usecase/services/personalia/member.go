package personalia

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	personaliaEntity "github.com/fiber-go-pos-app/internal/entity/personalia"

	personaliaRepo "github.com/fiber-go-pos-app/internal/repo/personalia"
)

func GetAllMember(ctx *fiber.Ctx) ([]personaliaEntity.Member, error) {
	members, err := personaliaRepo.GetAllMemberRepo(ctx)
	if err != nil {
		return []personaliaEntity.Member{}, err
	}
	return members, nil
}

func GetMemberByID(ctx *fiber.Ctx, ID string) (personaliaEntity.Member, error) {
	member, found, err := personaliaRepo.GetMemberByID(ctx, ID)
	if err != nil {
		return personaliaEntity.Member{}, err
	}

	if !found {
		return personaliaEntity.Member{}, fmt.Errorf("member dengan nama : %s tidak ditemukan", member.Name)
	}

	return member, nil
}

func InsertMember(ctx *fiber.Ctx, member personaliaEntity.Member) error {
	if err := personaliaRepo.InsertMember(ctx, member); err != nil {
		return err
	}

	return nil
}

func UpdateMember(ctx *fiber.Ctx, member personaliaEntity.Member) error {
	if _, err := GetMemberByID(ctx, member.ID); err != nil {
		return err
	}

	if err := personaliaRepo.UpdateMember(ctx, member); err != nil {
		return err
	}

	return nil
}

func DeleteMember(ctx *fiber.Ctx, ID string) error {
	if _, err := GetMemberByID(ctx, ID); err != nil {
		return err
	}

	if err := personaliaRepo.DeleteMember(ctx, ID); err != nil {
		return err
	}

	return nil
}

func UpsertMember(ctx *fiber.Ctx, member personaliaEntity.Member) error {
	_, found, err := personaliaRepo.GetMemberByID(ctx, member.ID)
	if err != nil {
		return err
	}

	if !found {
		return personaliaRepo.InsertMember(ctx, member)
	} else {
		return personaliaRepo.UpdateMember(ctx, member)
	}
}
