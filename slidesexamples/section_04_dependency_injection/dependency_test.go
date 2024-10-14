package dependency

import (
	"dependencyinjection/pkg/users"
	"errors"
	"testing"
)

var returnOneUser = func() ([]users.User, error) {
	return []users.User{{"foo", 12}}, nil
}
var failToGet = func() ([]users.User, error) {
	return nil, errors.New("failed")
}

func TestAverageAgeReplace(t *testing.T) {
	old := usersGet
	t.Run("oneUser", func(t *testing.T) {
		usersGet = returnOneUser
		t.Cleanup(func() { usersGet = old })
		avg, _ := UsersAverageAgeReplace()
		if avg != 12 {
			t.Fail()
		}
	})
	t.Run("with Err", func(t *testing.T) {
		usersGet = failToGet
		t.Cleanup(func() { usersGet = old })
		_, err := UsersAverageAgeReplace()
		if err == nil {
			t.Fail()
		}
	})
}

func TestAverageAgeInj(t *testing.T) {
	t.Run("oneUser", func(t *testing.T) {
		avg, _ := UsersAverageAgeInj(returnOneUser)
		if avg != 12 {
			t.Fail()
		}
	})
	t.Run("with Err", func(t *testing.T) {
		_, err := UsersAverageAgeInj(failToGet)
		if err == nil {
			t.Fail()
		}
	})
}
