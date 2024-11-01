package dependency

import "dependencyinjection/pkg/users"

func AppUsersAverageAge() (int, error) {
	app := users.NewApplication()
	users, err := app.Users()
	if err != nil {
		return 0, err
	}
	return averageAgeForUsers(users), nil
}

type UsersGetter interface {
	Users() ([]users.User, error)
}

var _ UsersGetter = users.Application{}

func AppUsersAverageAgeInj(getter UsersGetter) (int, error) {
	users, err := getter.Users()
	if err != nil {
		return 0, err
	}
	return averageAgeForUsers(users), nil
}
