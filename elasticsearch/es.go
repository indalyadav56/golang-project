package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Inserting a document data into elasticsearch
	doc := map[string]interface{}{
		"title":   "Sample Document",
		"content": "This is a sample document for Elasticsearch with Golang.",
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf("Error encoding document: %s", err)
	}

	// Indexing the document
	index := "my_index"
	_, err = es.Index(index, &buf)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}

	// Searching for a document
	var (
		r  map[string]interface{}
		esResponse *esapi.Response
	)

	// Prepare a query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "sample",
			},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request
	esResponse, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(index),
		es.Search.WithBody(strings.NewReader(string(queryJSON))),
	)
	if err != nil {
		log.Fatalf("Error searching for documents: %s", err)
	}
	defer esResponse.Body.Close()

	if err := json.NewDecoder(esResponse.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Display search results
	fmt.Printf("Search Results:\n%+v\n", r)
}
