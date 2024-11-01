package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func reply(path string) []byte {
	expected, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", path))
	if err != nil {
		panic("path not found")
	}
	return expected
}

func replyHelper(path string, t *testing.T) []byte {
	t.Helper()
	expected, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", path))
	if err != nil {
		panic("path not found")
	}
	return expected
}

func TestFetchUsers(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := reply(r.URL.Path)
		w.Write(res)
	}))
	t.Cleanup(svr.Close)

	toCheck, err := FetchUsers(svr.URL + "/basic")
	if err != nil {
		t.Error("received error", err)
	}
	if len(toCheck) != 1 {
		t.Fail()
	}
}
