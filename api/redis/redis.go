package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetClient() *redis.Client {
	return client
}

func ExampleClient() {

	fmt.Println(client)

	// get a key-value pair
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func Set(prefix, key, value any) {
	// set a key-value pair
	k := fmt.Sprintf("%s.%s", prefix, key)
	err := client.Set(k, value, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Redis variable has been set [%s : %s]\n", k, value)
}

func Get(prefix, key string) string {
	k := fmt.Sprintf("%s.%s", prefix, key)
	val, err := client.Get(k).Result()
	if err != nil {
		panic(err)
	}

	return val
}

func Delete(prefix, key string) (int64, error) {
	k := fmt.Sprintf("%s.%s", prefix, key)
	val, err := client.Del(k).Result()
	if err != nil {
		return 0, err

	}

	return val, nil
}

func Keys(prefix, key string) []string {
	k := fmt.Sprintf("%s.%s", prefix, key)
	val, err := client.Keys(k).Result()
	if err != nil {
		panic(err)
	}

	return val
}

func GetValues(prefix string) map[string]string {
	keys := Keys(prefix, "*")
	vals := map[string]string{}
	for _, key := range keys {
		val, err := client.Get(key).Result()
		if err != nil {
			panic(err)
		}
		vals[key] = val
	}

	return vals
}

func Incr(prefix, key string) {
	k := fmt.Sprintf("%s.%s", prefix, key)
	_, err := client.Incr(k).Result()
	if err != nil {
		panic(err)
	}
}
