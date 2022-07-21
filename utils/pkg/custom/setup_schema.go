package custom

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gofiber/fiber/v2"
	"github.com/tanimutomo/sqlfile"

	constantsEntity "github.com/fiber-go-pos-app/internal/entity/constants"

	"github.com/fiber-go-pos-app/utils/pkg/databases/elasticsearch"
	"github.com/fiber-go-pos-app/utils/pkg/databases/postgres"
)

func SetupSchema() error {
	if err := SetupElasticSearch(); err != nil {
		return err
	}

	if err := SetupPostgresTable(); err != nil {
		return err
	}

	return nil
}

func SetupElasticSearch() error {
	var indices []string
	var mappings []string
	wg := new(sync.WaitGroup)

	esClient := elasticsearch.GetESClient()

	files, err := ioutil.ReadDir(constantsEntity.ElasticSearchSchemeDirectory)
	if err != nil {
		return err
	}

	path, err := filepath.Abs(constantsEntity.ElasticSearchSchemeDirectory)
	if err != nil {
		return err
	}

	for _, f := range files {
		// Initialization filename and body of file (mapping)
		fileName := ReadFileNameWithoutExtension(f.Name())
		body, _ := ioutil.ReadFile(filepath.Join(path, f.Name()))

		// Append of index
		indices = append(indices, fileName)
		mappings = append(mappings, string(body))
	}

	response, err := esClient.Indices.Exists(indices)
	if err != nil {
		return err
	}

	if response.StatusCode != fiber.StatusNotFound {
		return nil
	}

	errorCreateIndex := make(chan error)
	for idx, elasticIndex := range indices {
		wg.Add(1)

		// defined new variable to handle race condition on go routines
		newIdx := idx
		newElasticIndex := elasticIndex

		go func() {
			err := func(newIdx int, newElasticIndex string, wg *sync.WaitGroup) error {
				defer wg.Done()

				indexReq := esapi.IndicesCreateRequest{
					Index: newElasticIndex,
					Body:  strings.NewReader(mappings[newIdx]),
				}

				response, err := indexReq.Do(context.Background(), esClient)
				if err != nil {
					errorCreateIndex <- err
					return err
				}

				if response.IsError() {
					return fmt.Errorf("error create index : %s", newElasticIndex)
				}

				return nil
			}(newIdx, newElasticIndex, wg)

			if err != nil {
				errorCreateIndex <- err
				return
			}

		}()
	}
	wg.Wait()

	select {
	case err, ok := <-errorCreateIndex:
		if ok {
			return err
		}
	default:
	}

	return nil
}

func SetupPostgresTable() error {
	// Initialize Connection
	sqlFile := sqlfile.New()
	db := postgres.GetPgConn()

	// Load schema folder
	if err := sqlFile.Directory(constantsEntity.PostgresSchemeDirectory); err != nil {
		return err
	}

	// Execute the stored queries
	// transaction is used to execute queries in Exec()
	_, err := sqlFile.Exec(db.DB)
	return err
}
