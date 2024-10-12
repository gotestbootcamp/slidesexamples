package parse

import (
	"encoding/json"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Parse(fileName string) (User, error) {
	res := User{}

	b, err := os.ReadFile(fileName)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func ParseAndIncrementAge(fileName string) (User, error) {
	parsed, err := Parse(fileName)
	if err != nil {
		return User{}, err
	}
	parsed.Age++
	return parsed, nil
}
