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

func ParseWithReader(fileName string) (User, error) {
	res := User{}

	f, err := os.Open(fileName)
	if err != nil {
		return User{}, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil

}
