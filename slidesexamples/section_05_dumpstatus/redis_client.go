package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"testcontainer/users"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	rdb *redis.Client
}

func NewStorage(url string) *Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Storage{rdb: rdb}
}

func (s *Storage) AddUser(ctx context.Context, user users.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// I know this is racy, trying to make a point
	_, err = s.rdb.Get(ctx, user.Name).Result()
	if err != redis.Nil {
		return fmt.Errorf("user already exists")
	}

	_, err = s.rdb.Set(ctx, user.Name, string(jsonUser), 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUser(ctx context.Context, name string) (users.User, error) {
	res := users.User{}

	userJson, err := s.rdb.Get(ctx, name).Result()
	if err == redis.Nil {
		return res, fmt.Errorf("user %s does not exist", name)
	}
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(userJson), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
