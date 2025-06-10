package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	// Initialize GORM and create SQLite database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	db.AutoMigrate(&Product{})

	// Initialize Elasticsearch client
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", 
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Create a product
	r.POST("/products", func(c *gin.Context) {
		var product Product
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Save product to SQLite database
		db.Create(&product)

		// Index product in Elasticsearch
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode product"})
			return
		}

		_, err := es.Index("products", &buf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to index product in Elasticsearch"})
			return
		}

		c.JSON(http.StatusCreated, product)
	})

	// Search for products in Elasticsearch
	r.GET("/products", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Search query parameter 'q' is required"})
			return
		}

		var (
			r  map[string]interface{}
			esResponse *esapi.Response
		)

		// Prepare the Elasticsearch query
		searchQuery := map[string]interface{}{
			"query": map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":  query,
					"fields": []string{"name", "description"},
				},
			},
		}

		searchQueryJSON, err := json.Marshal(searchQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode search query"})
			return
		}

		// Perform the search request
		esResponse, err = es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("products"),
			es.Search.WithBody(strings.NewReader(string(searchQueryJSON))),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products in Elasticsearch"})
			return
		}
		defer esResponse.Body.Close()

		if err := json.NewDecoder(esResponse.Body).Decode(&r); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse search response"})
			return
		}

		c.JSON(http.StatusOK, r)
	})

	// Run Gin server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %s", err)
	}
}
