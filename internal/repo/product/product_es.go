package product

import (
	"fmt"
	elasticsearch2 "github.com/fiber-go-pos-app/utils/pkg/databases/elasticsearch"

	"github.com/gofiber/fiber/v2"

	productEntity "github.com/fiber-go-pos-app/internal/entity/product"
)

const queryGetProductES = `{
	"from" : %d,
    "size" : %d,
	"sort" : [
		{"product_id" : "asc" }
	],
    "query": {
        "bool": {
			"should": [
				{
		            "wildcard": {
		                "product_id": "*%[3]s*"
		            }
		        },
				{
		            "wildcard": {
		                "name": "*%[3]s*"
		            }
		        },
				{
		            "wildcard": {
		                "barcode": "*%[3]s*"
		            }
		        }
			]
        }
    }
}
`

type resGetProductES struct {
	Hits struct {
		Hits []struct {
			Source productEntity.Product `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func GetProductListES(ctx *fiber.Ctx, offset int, limit int, search string) ([]productEntity.Product, error) {
	var destination resGetProductES
	query := fmt.Sprintf(queryGetProductES, offset, limit, search)
	err := elasticsearch2.Search(ctx.Context(), productEntity.ESProductIndex, query, &destination)
	return buildESToProductList(destination), err
}

func buildESToProductList(data resGetProductES) []productEntity.Product {
	products := make([]productEntity.Product, len(data.Hits.Hits))
	for index, product := range data.Hits.Hits {
		products[index] = product.Source
	}
	return products
}

const queryGetCountProductES = `{
    "query": {
        "bool": {
			"must": [
				{
		            "wildcard": {
		                "name": "*%s*"
		            }
		        }
			]
        }
    }
}
`

type resGetCountProductES struct {
	Count int64 `json:"count"`
}

func GetCountProductES(ctx *fiber.Ctx, search string) (int64, error) {
	var destination resGetCountProductES
	query := fmt.Sprintf(queryGetCountProductES, search)
	err := elasticsearch2.Count(ctx.Context(), productEntity.ESProductIndex, query, &destination)
	return destination.Count, err
}

func UpsertProductES(ctx *fiber.Ctx, product productEntity.Product) error {
	return elasticsearch2.Upsert(ctx.Context(), productEntity.ESProductIndex, product.ProductID, product)
}

func DeleteProductES(ctx *fiber.Ctx, productID string) error {
	return elasticsearch2.Delete(ctx.Context(), productEntity.ESProductIndex, productID)
}
