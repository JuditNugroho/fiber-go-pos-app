package elasticsearch

import "github.com/elastic/go-elasticsearch/v7"

var esClient *elasticsearch.Client

func NewESClient() error {
	var err error
	esClient, err = elasticsearch.NewDefaultClient()
	return err
}

func GetESClient() *elasticsearch.Client {
	return esClient
}
