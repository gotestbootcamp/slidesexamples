package redisclient

import (
	"context"
	"testcontainer/users"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

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

	storage := NewStorage("127.0.0.1:" + mapped.Port())

	err := storage.AddUser(context.Background(), users.User{"foo", 12})
	if err != nil {
		t.Fatal("add user failed", err)
	}
	user, err := storage.GetUser(context.Background(), "foo")
	if err != nil {
		t.Fatal("get user failed", err)
	}
	if user.Age != 12 {
		t.Fatal("age is not 12")
	}
}

type fakeRedisClient struct {
	lastCall []interface{}
}

func (f *fakeRedisClient) Do(ctx context.Context, cmd ...interface{}) *redis.Cmd {
	f.lastCall = cmd
	return &redis.Cmd{}
}

func (f *fakeRedisClient) Get(context.Context, string) *redis.StringCmd {
	return nil
}

func TestWithMock(t *testing.T) {
	f := &fakeRedisClient{}
	s := Storage{client: f}

	s.AddUser(context.TODO(), users.User{"foo", 12})

	if f.lastCall[0].(string) != "set" {
		t.Fatal()
	}
}
