package benchmark_redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"testing"
	"time"
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
	cnt := getClient()
	params := []int{10, 20, 50, 100, 200, 1000, 5000}
	for _, item := range params {
		cnt.Set(ctx, fmt.Sprintf("key_%d", item), getRandomString(item), 0)
	}
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func BenchmarkRedisKv_10b(b *testing.B) {
	//b.Log("start to test for 10 bytes value...")
	getRedisKv(b, "key_10")
}

func BenchmarkRedisKv_20b(b *testing.B) {
	//b.Log("start to test for 20 bytes value...")
	getRedisKv(b, "key_20")
}

func BenchmarkRedisKv_50b(b *testing.B) {
	//b.Log("start to test for 50 bytes value...")
	getRedisKv(b, "key_50")
}

func BenchmarkRedisKv_100b(b *testing.B) {
	//b.Log("start to test for 100 bytes value...")
	getRedisKv(b, "key_100")
}

func BenchmarkRedisKv_200b(b *testing.B) {
	//b.Log("start to test for 200 bytes value...")
	getRedisKv(b, "key_200")
}

func BenchmarkRedisKv_1k(b *testing.B) {
	//b.Log("start to test for 1 kb value...")
	getRedisKv(b, "key_1000")
}

func BenchmarkRedisKv_5k(b *testing.B) {
	//b.Log("start to test for 5 kb value...")
	getRedisKv(b, "key_5000")
}

func getRedisKv(b *testing.B, key string) {
	b.StopTimer()
	client := getClient()
	b.StartTimer()
	for i := 1; i < b.N; i++ {
		_ = client.Get(ctx, key).Val()
		//b.Log(val)
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

func TestCalcAvgMemory(t *testing.T) {
	cnt := getClient()
	params := []int{10, 20, 50, 100, 200, 1000, 5000}
	for _, item := range params {
		cnt.FlushDB(ctx)
		for i := 0; i < 10000; i++ {
			cnt.Set(ctx, fmt.Sprintf("key_%d", i), getRandomString(item), 0)
		}
	}
}
