package elasticsearch

import (
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func Delete(ctx context.Context, index string, documentID string) error {
	req := esapi.DeleteRequest{
		Index:      index,
		DocumentID: documentID,
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(ctx, esClient)
	if err != nil {
		return fmt.Errorf("error getting response: %s", err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error delete es : %s", res.Status())
	}
	return nil
}
