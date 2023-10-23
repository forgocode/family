package redis

import (
	"fmt"
	"testing"

	"github.com/forgocode/family/internal/conf"
)

func TestInitRedis(t *testing.T) {
	config := conf.GetConfig()
	_, err := InitRedis(fmt.Sprintf("%s:%d", config.Redis.IP, config.Redis.Port), config.Redis.Password)
	if err != nil {
		t.Logf("failed to connect to redis, err: %+v\n", err)
		return
	}
	t.Log("successfully to connect to redis")
}
