package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", MyCustomHandler)

	pprof.Register(router)

	router.Run(":8080")
}

func MyCustomHandler(c *gin.Context) {
	HeavyComputation()
	c.String(200, "This is a custom handler")
}

func HeavyComputation() {
	for i := 0; i < 100_00_000; i++ {
		_ = i * i // Simulating some work
	}
}
