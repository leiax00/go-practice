package benchmark_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var rdb *redis.Client
var ctx = context.Background()

func TestMain(m *testing.M) {
	fmt.Println("begin")
	initData()
	m.Run()
	fmt.Println("end")
}

func initData() {
	//cnt := getClient()
	//cnt.Set(ctx, )
}

func BenchmarkRedisKv_10b(b *testing.B) {
	b.Log("start to test for 10 bytes value...")
	getRedisKv(b, "key_10b")
}

func BenchmarkRedisKv_20b(b *testing.B) {
	b.Log("start to test for 20 bytes value...")
	getRedisKv(b, "key_20b")
}

func BenchmarkRedisKv_50b(b *testing.B) {
	b.Log("start to test for 50 bytes value...")
	getRedisKv(b, "key_50b")
}

func BenchmarkRedisKv_100b(b *testing.B) {
	b.Log("start to test for 100 bytes value...")
	getRedisKv(b, "key_100b")
}

func BenchmarkRedisKv_200b(b *testing.B) {
	b.Log("start to test for 200 bytes value...")
	getRedisKv(b, "key_200b")
}

func BenchmarkRedisKv_1k(b *testing.B) {
	b.Log("start to test for 1 kb value...")
	getRedisKv(b, "key_1k")
}

func BenchmarkRedisKv_5k(b *testing.B) {
	b.Log("start to test for 5 kb value...")
	getRedisKv(b, "key_5k")
}

func getRedisKv(b *testing.B, key string) {
	b.StopTimer()
	client := getClient()
	b.StartTimer()
	for i := 1; i < b.N; i++ {
		val := client.Get(ctx, key).Val()
		b.Log(val)
	}
}

func getClient() *redis.Client {
	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

	}
	return rdb
}
