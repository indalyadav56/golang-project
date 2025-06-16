package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUserHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/user", nil)

	CreateUserHandler(w, r)
	fmt.Println(w.Result().Status)
	if w.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, w.Code)
	}
	if w.Body.String() != "user created successfully" {
		t.Errorf("expected body %s, got %s", "user created", w.Body.String())
	}
}

func BenchmarkCreateUserHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/user", nil)
		CreateUserHandler(w, r)
	}
}
