package product

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"
	productEntity "github.com/fiber-go-pos-app/internal/entity/product"
	customPkg "github.com/fiber-go-pos-app/utils/pkg/custom"

	productRepo "github.com/fiber-go-pos-app/internal/repo/product"
)

// GetDTAllProduct : Get List Of Product for Datatable
func GetDTAllProduct(ctx *fiber.Ctx, page int, limit int, search string) (productEntity.ListProductDataResponse, error) {
	esLimit, offset := customPkg.BuildOffsetAndLimitES(page, limit)

	products, err := productRepo.GetProductListES(ctx, offset, esLimit, search)
	if err != nil {
		return productEntity.ListProductDataResponse{}, err
	}

	totalProduct, err := productRepo.GetCountProductES(ctx, search)
	if err != nil {
		return productEntity.ListProductDataResponse{}, err
	}

	return productEntity.ListProductDataResponse{
		Total: totalProduct,
		Data:  products,
	}, nil
}

func GetProductByID(ctx *fiber.Ctx, ID string) (productEntity.Product, error) {
	product, found, err := productRepo.GetProductByID(ctx, ID)
	if err != nil {
		return productEntity.Product{}, err
	}

	if !found {
		return productEntity.Product{}, fmt.Errorf("product dengan nama : %s tidak ditemukan", product.Name)
	}

	return product, nil
}

func GetProductByIDOrBarcode(ctx *fiber.Ctx, search string) (productEntity.Product, error) {
	product, found, err := productRepo.GetProductByIDOrBarcode(ctx, search)
	if err != nil {
		return productEntity.Product{}, err
	}

	if !found {
		return productEntity.Product{}, errors.New(constantsEntity.ErrNoDataFound)
	}

	return product, nil
}

func InsertProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	if err := productRepo.InsertProduct(ctx, product); err != nil {
		return err
	}

	_ = productRepo.UpsertProductES(ctx, product)

	return nil
}

func UpdateProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	if _, err := GetProductByID(ctx, product.ProductID); err != nil {
		return err
	}

	if err := productRepo.UpdateProduct(ctx, product); err != nil {
		return err
	}

	_ = productRepo.UpsertProductES(ctx, product)

	return nil
}

func DeleteProduct(ctx *fiber.Ctx, productID string) error {
	if _, err := GetProductByID(ctx, productID); err != nil {
		return err
	}

	if err := productRepo.DeleteProduct(ctx, productID); err != nil {
		return err
	}

	_ = productRepo.DeleteProductES(ctx, productID)

	return nil
}

func UpsertProduct(ctx *fiber.Ctx, product productEntity.Product) error {
	var err error
	_, found, err := productRepo.GetProductByID(ctx, product.ProductID)
	if err != nil {
		return err
	}

	if !found {
		err = productRepo.InsertProduct(ctx, product)
	} else {
		err = productRepo.UpdateProduct(ctx, product)
	}

	if err != nil {
		return err
	}

	_ = productRepo.UpsertProductES(ctx, product)

	return nil
}
