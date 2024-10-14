package dependency

import (
	"dependencyinjection/pkg/users"
	"errors"
	"testing"
)

type mockApp struct {
	called    int
	usersRes  []users.User
	shouldErr bool
}

func (m *mockApp) Users() ([]users.User, error) {
	m.called++
	if m.shouldErr {
		return nil, errors.New("failed")
	}
	return m.usersRes, nil
}

var _ UsersGetter = &mockApp{}

func TestSuccess(t *testing.T) {
	m := &mockApp{
		called:    0,
		usersRes:  []users.User{{"foo", 12}, {"bar", 14}},
		shouldErr: false}
	res, err := AppUsersAverageAgeInj(m)
	if m.called != 1 {
		t.Fail()
	}
	if res != 13 {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestFail(t *testing.T) {
	m := &mockApp{
		called:    0,
		usersRes:  nil,
		shouldErr: true}
	_, err := AppUsersAverageAgeInj(m)
	if m.called != 1 {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}
