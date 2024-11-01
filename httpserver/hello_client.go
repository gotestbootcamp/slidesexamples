package hello

import (
	"encoding/json"
	"fmt"
	"httptest/users"
	"net/http"
)

func FetchUsers(url string) ([]users.User, error) {
	httpRes, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to complete Get request %w", err)
	}

	res := []users.User{}
	err = json.NewDecoder(httpRes.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
