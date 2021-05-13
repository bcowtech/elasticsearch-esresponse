package esresponse

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	esresponse "github.com/bcowtech/elasticsearch-esresponse"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

func TestSearch(t *testing.T) {
	setupSearchTestCase()

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

	req := esapi.SearchRequest{
		Index: []string{"gaasdemo-service-event-log-20201228-724001"},
	}

	resp, err := req.Do(context.Background(), client)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Error("should not be error")
	}

	result, err := esresponse.AsSearchResult(resp.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		t.Error("should not be error")
	}

	// fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", result)
	resp.Body.Close()
}

func setupSearchTestCase() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_ADDRESS"),
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	req := esapi.IndexRequest{
		Index:      "gaasdemo-service-event-log-20201228-724001",
		DocumentID: "192.168.56.54#0001",
		Body:       bytes.NewReader(getSearchTestCase()),
		Refresh:    "true",
	}

	resp, err := req.Do(context.Background(), client)
	if err != nil {
		panic(err)
	}

	if resp.IsError() {
		panic(fmt.Errorf("fail get status %s", resp.Status()))
	}
}

func getSearchTestCase() []byte {
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
