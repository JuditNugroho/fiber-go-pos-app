package product

import (
	"database/sql"
	"github.com/fiber-go-pos-app/utils/pkg/databases/postgres"

	"github.com/gofiber/fiber/v2"

	productEntity "github.com/fiber-go-pos-app/internal/entity/product"
)

const queryGetProductByID = `
	SELECT product_id, name, barcode, stock, ppn, price, member_price, discount, category_id
	FROM products
	WHERE product_id = $1
`

func GetProductByID(ctx *fiber.Ctx, productID string) (productEntity.Product, bool, error) {
	var product productEntity.Product
	db := postgres.GetPgConn()

	if err := db.GetContext(ctx.Context(), &product, queryGetProductByID, productID); err != nil {
		if err == sql.ErrNoRows {
			return product, false, nil
		}
		return product, false, err
	}
	return product, true, nil
}

const queryGetProductByIDOrBarcode = `
	(SELECT product_id, name, barcode, stock, ppn, price, member_price, discount, category_id
	FROM products
	WHERE product_id = $1)
	UNION
	(SELECT product_id, name, barcode, stock, ppn, price, member_price, discount, category_id
	FROM products
	WHERE barcode = $1)
`

func GetProductByIDOrBarcode(ctx *fiber.Ctx, search string) (productEntity.Product, bool, error) {
	var product productEntity.Product
	db := postgres.GetPgConn()

	if err := db.GetContext(ctx.Context(), &product, queryGetProductByIDOrBarcode, search); err != nil {
		if err == sql.ErrNoRows {
			return product, false, nil
		}
		return product, false, err
	}
	return product, true, nil
}

const insertProduct = `
	INSERT INTO products (product_id, name, barcode, stock, ppn, price, member_price, discount, category_id)
	VALUES (:product_id, :name, :barcode, :stock, :ppn, :price, :member_price, :discount, :category_id)
`

func InsertProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	db := postgres.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), insertProduct, product)
	if err != nil {
		return err
	}
	return nil
}

const updateProduct = `
	UPDATE products SET
		name = :name,
	    barcode = :barcode,
	    stock = :stock,
	    ppn = :ppn,
	    price = :price,
	    member_price = :member_price,
	    discount = :discount,
	    category_id = :category_id,
		update_time = NOW()
	WHERE product_id = :product_id
`

func UpdateProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	db := postgres.GetPgConn()

	_, err := db.NamedQueryContext(ctx.Context(), updateProduct, product)
	if err != nil {
		return err
	}
	return nil
}

const deleteProduct = `
	DELETE FROM products
	WHERE product_id = $1
`

func DeleteProduct(ctx *fiber.Ctx, productID string) error {
	db := postgres.GetPgConn()

	_, err := db.ExecContext(ctx.Context(), deleteProduct, productID)
	if err != nil {
		return err
	}
	return nil
}
