package elasticsearch

import (
	"bytes"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	goccyJson "github.com/goccy/go-json"
)

func Upsert(ctx context.Context, index string, documentID string, data any) error {
	// Build the request body.
	body, err := goccyJson.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling document: %s", err)
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		Index:      index,
		Refresh:    "true",
		DocumentID: documentID,
		Body:       bytes.NewReader(body),
	}

	// Perform the request with the client.
	res, err := req.Do(ctx, esClient)
	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error insert es : %s", res.Status())
	}
	return nil
}
