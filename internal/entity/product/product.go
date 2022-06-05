package product

type Product struct {
	ProductID   string  `json:"product_id" db:"product_id" validate:"required,max=20"`
	Name        string  `json:"name" db:"name" validate:"max=255"`
	Barcode     string  `json:"barcode" db:"barcode"`
	Stock       int64   `json:"stock" db:"stock" validate:"min=0"`
	Ppn         bool    `json:"ppn" db:"ppn"`
	Price       float64 `json:"price" db:"price" validate:"min=0"`
	MemberPrice float64 `json:"member_price" db:"member_price" validate:"min=0"`
	Discount    float64 `json:"discount" db:"discount"`
	CategoryId  int     `json:"category_id" db:"category_id"`
}

const ESProductIndex = "products"

type ListProductDataResponse struct {
	Total int64
	Data  []Product
}
