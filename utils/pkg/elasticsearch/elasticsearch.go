package elasticsearch

import "github.com/elastic/go-elasticsearch/v7"

var ESClient *elasticsearch.Client

func NewESClient() error {
	var err error
	ESClient, err = elasticsearch.NewDefaultClient()
	return err
}
