package dependency

import (
	"dependencyinjection/pkg/users"
	"encoding/json"
	"io"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func parseReader(r io.Reader) (User, error) {
	res := User{}
	err := json.NewDecoder(r).Decode(&res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func Parse(fileName string) (User, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return User{}, err
	}
	defer f.Close()
	return parseReader(f)
}

func UsersAverageAge() (int, error) {
	users, err := users.Get()
	if err != nil {
		return 0, err
	}
	return averageAgeForUsers(users), nil
}

var usersGet = users.Get

func UsersAverageAgeReplace() (int, error) {
	users, err := usersGet()
	if err != nil {
		return 0, err
	}
	return averageAgeForUsers(users), nil
}

type usersRetriever func() ([]users.User, error)

func UsersAverageAgeInj(findUsers usersRetriever) (int, error) {
	users, err := findUsers()
	if err != nil {
		return 0, err
	}
	return averageAgeForUsers(users), nil
}

func averageAgeForUsers(users []users.User) int {
	sum := 0
	for _, u := range users {
		sum += u.Age
	}
	return sum / len(users)
}
