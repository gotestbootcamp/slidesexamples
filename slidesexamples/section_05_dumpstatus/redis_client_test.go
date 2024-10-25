package redisclient

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testcontainer/users"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func dumpRedisContent(url string, rdb *redis.Client, t *testing.T) string {
	t.Helper()
	ctx := context.Background()

	iter := rdb.Scan(ctx, 0, "", 0).Iterator()
	res := ""
	for iter.Next(ctx) {
		key := iter.Val()
		val, _ := rdb.Get(ctx, key).Result()
		res = res + fmt.Sprintf("%s: %s\n", key, val)
	}
	filename := strings.Replace(t.Name(), "/", "-", -1) + ".dump"
	if err := os.WriteFile(filename, []byte(res), 0666); err != nil {
		t.Fatal(err)
	}
	return res
}

func TestWithRedis(t *testing.T) {
	if testing.Short() {
		t.Skip("container test, skipping with -short")
	}
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, _ := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	mapped, _ := redisC.MappedPort(context.Background(), "6379/tcp")

	t.Cleanup(func() { redisC.Terminate(context.Background()) })

	url := "127.0.0.1:" + mapped.Port()
	storage := NewStorage(url)
	t.Run("add", func(t *testing.T) {
		err := storage.AddUser(context.Background(), users.User{"foo", 12})
		if err != nil {
			dumpRedisContent(url, t)
			t.Fatal("add user failed", err)
		}
	})

	t.Run("add and get", func(t *testing.T) {
		err := storage.AddUser(context.Background(), users.User{"foo", 12})
		if err != nil {
			dumpRedisContent(url, t)
			t.Fatal("add user failed", err)
		}
		user, err := storage.GetUser(context.Background(), "foo")
		if err != nil {
			dumpRedisContent(url, t)
			t.Fatal("get user failed", err)
		}
		if user.Age != 12 {
			dumpRedisContent(url, t)
			t.Fatal("age is not 12")
		}
	})
}
