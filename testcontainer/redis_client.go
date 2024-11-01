package redisclient

import (
	"context"
	"encoding/json"
	"fmt"
	"testcontainer/users"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	client redisClient
}

type redisClient interface {
	Do(context.Context, ...interface{}) *redis.Cmd
	Get(context.Context, string) *redis.StringCmd
}

func NewStorage(url string) *Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Storage{client: rdb}
}

func (s *Storage) AddUser(ctx context.Context, user users.User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, "sel", user.Name, string(jsonUser)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUser(ctx context.Context, name string) (users.User, error) {
	res := users.User{}

	userJson, err := s.client.Get(ctx, name).Result()
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
