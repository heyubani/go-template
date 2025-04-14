package connections

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/heyubani/go-template/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var redisOnce sync.Once

// This creates a redis client that has tracing enabled using Datadog
func createRedisConnection() *redis.Client {
	opt := &redis.Options{
		Addr:        config.AppConfig.Redis.Host,
		Password:    config.AppConfig.Redis.Password,
		DialTimeout: time.Second * 20,
		DB:          config.AppConfig.Redis.Db,
	}

	client := redis.NewClient(opt)

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("cannot connext to redis. Error: %s, pong message: %#v ", err.Error(), pong)
		panic("cannot connect to redis. Error: " + err.Error())
	}

	return client
}

func GetRedisConnection() *redis.Client {
	redisOnce.Do(func() {
		redisClient = createRedisConnection()
	})
	return redisClient
}

// SetRedisKey set a redis key and value to the application redis instance
func SetRedisKey(key string, value interface{}, expiration time.Duration) (valid bool, result interface{}) {
	result, err := redisClient.Set(context.Background(), key, value, expiration).Result()

	if err != nil {
		if err != redis.Nil {
			log.Printf("redis error @SetRedisKey - %s", err.Error())
		}
		return false, nil
	}
	return true, result
}

func RetrieveGenericStruct(key string, dest interface{}) (bool, error) {
	value, err := redisClient.Get(context.Background(), key).Result()

	if err != nil {
		if err != redis.Nil {
			log.Printf("redis error @RetrieveGenericStruct - %s", err.Error())
		}
		return false, err
	}

	b := []byte(value)

	errJ := json.Unmarshal(b, dest)

	if errJ != nil {
		return false, errJ
	}

	return true, nil
}
