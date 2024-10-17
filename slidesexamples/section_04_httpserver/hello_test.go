package hello

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	router := newRouter()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	res := w.Result()

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != "<!DOCTYPE html>\nhello, Gopher!\n" {
		t.Errorf("Expected hello but got %v", string(data))
	}
}
