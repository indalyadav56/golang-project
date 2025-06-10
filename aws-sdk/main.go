package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func main() {
	// var bucket, key string
	// var timeout time.Duration

	// flag.StringVar(&bucket, "b", "", "Bucket name.")
	// flag.StringVar(&key, "k", "", "Object key name.")
	// flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	// flag.Parse()

	// fmt.Println("bucketName:->", bucket, key, timeout)

	awsConfig := aws.Config{
		Region: "us-west-2",
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(
			"YOUR_ACCESS_KEY_ID",
			"YOUR_SECRET_ACCESS_KEY",
			"",
		)),
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithLogConfigurationWarnings())
	// if err != nil {
	// 	log.Fatalf("unable to load SDK config, %v", err)
	// }

	// Create S3 service client
	client := s3.NewFromConfig(awsConfig)

	client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	// // List buckets
	// resp, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	// if err != nil {
	//     log.Fatalf("unable to list buckets, %v", err)
	// }

	// fmt.Println("Buckets:")
	// for _, bucket := range resp.Buckets {
	//     fmt.Printf("  %s\n", *bucket.Name)
	// }

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		fmt.Println("hello there")
		c.JSON(200, map[string]interface{}{
			"message": "hello there",
			"status":  "success",
		})
	})

	router.Run(":8080")

}
