package custom

import (
	"net/http"
	"strings"

	"github.com/tanimutomo/sqlfile"

	"github.com/fiber-go-pos-app/utils/pkg/databases/postgres"
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
	// Initialize Connection
	sqlFile := sqlfile.New()
	db := postgres.GetPgConn()

	// Load schema folder
	if err := sqlFile.Directory("./utils/schema/postgres"); err != nil {
		return err
	}

	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	_, err := sqlFile.Exec(db.DB)
	return err
}
