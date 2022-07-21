package elasticsearch

import (
	"bytes"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	goccyJson "github.com/goccy/go-json"
)

type SearchOption func(search esapi.Search) func(request *esapi.SearchRequest)

func SearchOptRouting(route ...string) SearchOption {
	return func(search esapi.Search) func(request *esapi.SearchRequest) {
		return search.WithRouting(route...)
	}
}

func Search(ctx context.Context, index string, query string, destination any, options ...SearchOption) error {
	buf := new(bytes.Buffer)
	buf.WriteString(query)

	searchOptions := []func(request *esapi.SearchRequest){
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(buf),
		esClient.Search.WithTrackTotalHits(true),
	}

	for _, opt := range options {
		searchOptions = append(searchOptions, opt(esClient.Search))
	}

	resp, err := esClient.Search(
		searchOptions...,
	)

	if err != nil {
		return fmt.Errorf("error getting response: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.IsError() {
		var errMapping map[string]interface{}
		if err := goccyJson.NewDecoder(resp.Body).Decode(&errMapping); err != nil {
			return fmt.Errorf("error parsing the response body: %s", err)
		} else {
			return fmt.Errorf("elasticsearch response error : [%d] %s:%s",
				resp.StatusCode,
				errMapping["error"].(map[string]interface{})["type"],
				errMapping["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := goccyJson.NewDecoder(resp.Body).Decode(&destination); err != nil {
		return fmt.Errorf("error parsing the response body: %s", err.Error())
	}

	return err
}
