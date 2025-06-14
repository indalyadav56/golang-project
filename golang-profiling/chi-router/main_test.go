package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserByIDHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)
	handleUserByID(w, r)
}

func BenchmarkTestCrud(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCRUDLoad()

	}
}
