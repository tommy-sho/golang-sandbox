package kvs

import (
	"testing"
	"time"

	redis_mock "github.com/tommy-sho/golang-sandbox/database/redis-unit-test/testhelper"
)

func TestSetToken(t *testing.T) {
	client := redis_mock.NewMockRedis(t)

	if err := SetToken(client, "test", 1); err != nil {
		t.Fatalf("unexpected error while SetToken '%#v'", err)
	}
	actual, err := client.Get("TOKEN_test").Result()
	if err != nil {
		t.Fatalf("unexpected error while get value %v", err)
	}

	if expected := "1"; expected != actual {
		t.Errorf("expedted value %s, actual value %v", expected, actual)
	}
}

func TestGetIDByToken(t *testing.T) {
	client := redis_mock.NewMockRedis(t)
	if err := client.Set("TOKEN_test", 1, time.Second*1000).Err(); err != nil {
		t.Fatalf("unexpected error while set test data %#v", err)
	}

	actual, err := GetIDByToken(client, "test")
	if err != nil {
		t.Fatalf("unexpected error while GetIDByToken '%#v'", err)
	}
	if expected := 1; expected != actual {
		t.Errorf("expected value '%#v', actual value '%#v'", expected, actual)
	}
}
