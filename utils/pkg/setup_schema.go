package pkg

import (
	"github.com/fiber-go-pos-app/utils/pkg/postgres"
	"net/http"
	"strings"
)

func SetupSchema() error {
	SetupElasticSearch()
	if err := SetupPostgresTable(); err != nil {
		return err
	}
	return nil
}

func SetupElasticSearch() {
	payloadBytes := `{
       "mappings": {
           "properties": {
               "product_id": { "type": "keyword" },
               "name": { "type": "text" },
               "barcode": { "type": "text" },
               "stock": { "type": "integer" },
               "ppn": { "type": "boolean" },
               "price": { "type": "float" },
               "member_price": { "type": "float" },
               "discount": { "type": "float" },
               "category_id": { "type": "integer" }
           }
       }
   }`
	body := strings.NewReader(payloadBytes)
	req, _ := http.NewRequest("PUT", "http://localhost:9200/products/", body)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
}

func SetupPostgresTable() error {
	db := postgres.GetPgConn()
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS members
		(
		    id VARCHAR(20) NOT NULL PRIMARY KEY,
		    name VARCHAR(255) NOT NULL,
		    phone VARCHAR(20) NOT NULL,
		    create_time timestamp not null default now(),
		    update_time timestamp
		);
		
		
		CREATE INDEX IF NOT EXISTS member_phone_idx ON members (phone);

		CREATE TABLE IF NOT EXISTS products
		(
		    product_id VARCHAR(20) NOT NULL PRIMARY KEY,
		    name VARCHAR(255) NOT NULL,
		    barcode VARCHAR(30) NOT NULL DEFAULT '0',
		    stock BIGINT NOT NULL DEFAULT 0,
		    ppn BOOLEAN NOT NULL DEFAULT FALSE ,
		    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
		    member_price NUMERIC(10, 2) NOT NULL DEFAULT 0,
		    discount NUMERIC(10, 2) NOT NULL DEFAULT 0,
		    category_id INTEGER NOT NULL DEFAULT 0,
		    create_time timestamp not null default now(),
		    update_time timestamp
		);
		
		
		CREATE INDEX IF NOT EXISTS barcode_products_idx ON products (barcode);
		
		CREATE TABLE IF NOT EXISTS users
		(
		    user_id     VARCHAR(20)  NOT NULL PRIMARY KEY,
		    user_name   VARCHAR(30)  NOT NULL,
		    full_name   VARCHAR(255) NOT NULL DEFAULT '',
		    password    VARCHAR(255) NOT NULL DEFAULT '',
		    is_admin    bool         NOT NULL DEFAULT false,
		    create_time timestamp    not null default now(),
		    update_time timestamp
		);
		
		CREATE UNIQUE INDEX IF NOT EXISTS user_is_admin ON users (user_id, is_admin);
	`)
	return err
}
