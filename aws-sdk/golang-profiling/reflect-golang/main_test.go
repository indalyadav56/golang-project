package main

import (
	"fmt"
	"testing"
)

func TestStructToMap(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		err, mapData := StructToMap("indal")
		if err != nil {
			t.Errorf("mapData should be nil")
		}

		fmt.Println("mapData", mapData)
	})
}
