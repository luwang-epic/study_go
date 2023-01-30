package main

import "github.com/go-redis/redis"

func getClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 100, // 连接池大小
	})
	return client
}


func putKeyValue(key string, value string) {
	client := getClient()
	client.Set(key, value, 0)
}

func getValue(key string) string {
	client := getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func main() {
	putKeyValue("test", "go_test")
	result := getValue("test")
	println(result)
}
