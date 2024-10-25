package users

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Get() ([]User, error) {
	return []User{
		{"foo", 12},
		{"bar", 13},
	}, nil
}

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a Application) Users() ([]User, error) {
	return []User{
		{"foo", 12},
		{"bar", 13},
	}, nil
}
