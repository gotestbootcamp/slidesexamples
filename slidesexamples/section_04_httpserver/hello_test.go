package hello

import (
	"encoding/json"
	"httptest/users"
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

func TestFetchUsers(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uu := []users.User{{"foo", 12}, {"bar", 13}}
		json.NewEncoder(w).Encode(uu)
	}))
	t.Cleanup(svr.Close)

	toCheck, err := FetchUsers(svr.URL)
	if err != nil {
		t.Error("received error", err)
	}
	if len(toCheck) != 2 {
		t.Fail()
	}
}
