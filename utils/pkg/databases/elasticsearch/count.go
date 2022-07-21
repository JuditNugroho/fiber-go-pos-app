package elasticsearch

import (
	"bytes"
	"context"
	"fmt"

	goccyJson "github.com/goccy/go-json"
)

func Count(ctx context.Context, index string, query string, destination any) error {
	buf := new(bytes.Buffer)
	buf.WriteString(query)

	resp, err := esClient.Count(
		esClient.Count.WithContext(ctx),
		esClient.Count.WithIndex(index),
		esClient.Count.WithBody(buf),
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
