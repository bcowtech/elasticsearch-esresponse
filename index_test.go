package esresponse_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
	esresponse "gitlab.bcowtech.de/bcow-go/elasticsearch-esresponse"
)

func TestIndex(t *testing.T) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_ADDRESS"),
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Error("should not be error")
	}

	req := esapi.IndexRequest{
		Index:      "gaasdemo-service-event-log-20201228-724001",
		DocumentID: "192.168.56.54#0001",
		Body:       bytes.NewReader(getIndexTestCase()),
		Refresh:    "true",
	}

	resp, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Error("should not be error")
	}

	result, err := esresponse.AsIndexResponse(resp.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Error("should not be error")
	}

	// fmt.Printf("%+v\n", resp)
	// fmt.Printf("%+v\n", result)
	resp.Body.Close()

	expectedIndex := "gaasdemo-service-event-log-20201228-724001"
	if result.Index != expectedIndex {
		t.Errorf("assert 'IndexResponse.Index':: excepted '%v', got '%v'", expectedIndex, result.Index)
	}
	expectedId := "192.168.56.54#0001"
	if result.Id != expectedId {
		t.Errorf("assert 'IndexResponse.Id':: excepted '%v', got '%v'", expectedId, result.Id)
	}
	{
		shards := result.Shards
		if shards == nil {
			t.Error("assert 'IndexResponse.shards':: should not nil")
		}
		if shards.Total < 1 {
			t.Errorf("assert 'IndexResponse.Total':: should greater than 0 but not %d", shards.Total)
		}
		if shards.Successful < 1 {
			t.Errorf("assert 'IndexResponse.Successful':: should greater than 0 but not %d", shards.Successful)
		}
	}
}

func getIndexTestCase() []byte {
	return []byte(`{
		"timestamp": 1560973500123,
		"event_id" : "192.168.56.54#0001",
		"category" : "WalletService",
		"source"   : "192.168.56.54",
		"type"     : "PASS",
		"version"  : "v0.1.0b",
		"message"  : "\"GET /downloads/product_2 HTTP/1.1\" 304 0 \"-\" \"Debian APT-HTTP/1.3 (0.9.7.9)\"",
		"details"  : {
			"request": {
				"method": "GET",
				"path": "/downloads/product_2",
				"query_string": "arg1=one&arg2=2&arg3=drei",
				"header": "Content-Type: application/json\r\nUser-Agent: Debian APT-HTTP/1.3 (0.9.7.9)\r\n",
				"body": "some request body"
			},
			"response" :{
				"status_code": 304,
				"header": "X-Response-Header: f/twXyy",
				"body": "some response body"
			}
		},
		"metric": {
			"elapsed_time": 122,
			"response_body_bytes": 138
		}
	}`)
}
