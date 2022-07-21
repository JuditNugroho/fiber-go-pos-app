package personalia

import (
	"database/sql"
	"github.com/fiber-go-pos-app/utils/pkg/databases/postgres"

	"github.com/gofiber/fiber/v2"

	memberEntity "github.com/fiber-go-pos-app/internal/entity/personalia"
)

const queryGetAllMember = `
	SELECT id, name, phone
	FROM members
	ORDER BY id
`

func GetAllMemberRepo(ctx *fiber.Ctx) ([]memberEntity.Member, error) {
	var members []memberEntity.Member
	db := postgres.GetPgConn()

	if err := db.SelectContext(ctx.Context(), &members, queryGetAllMember); err != nil {
		return members, err
	}
	return members, nil
}

const queryGetMemberByID = `
	SELECT id, name, phone
	FROM members
	WHERE id = $1
`

func GetMemberByID(ctx *fiber.Ctx, ID string) (memberEntity.Member, bool, error) {
	var member memberEntity.Member

	db := postgres.GetPgConn()

	if err := db.GetContext(ctx.Context(), &member, queryGetMemberByID, ID); err != nil {
		if err == sql.ErrNoRows {
			return member, false, nil
		}
		return member, false, err
	}
	return member, true, nil
}

const insertMember = `
	INSERT INTO members (id, name, phone)
	VALUES (:id, :name, :phone)
`

func InsertMember(ctx *fiber.Ctx, member memberEntity.Member) error {

	db := postgres.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), insertMember, member)
	if err != nil {
		return err
	}
	return nil
}

const updateMember = `
	UPDATE members SET
		name = :name,
	    phone = :phone,
		update_time = NOW()
	WHERE id = :id
`

func UpdateMember(ctx *fiber.Ctx, member memberEntity.Member) error {

	db := postgres.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateMember, member)
	if err != nil {
		return err
	}
	return nil
}

const deleteMember = `
	DELETE FROM members
	WHERE id = $1
`

func DeleteMember(ctx *fiber.Ctx, ID string) error {

	db := postgres.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteMember, ID)
	if err != nil {
		return err
	}
	return nil
}
