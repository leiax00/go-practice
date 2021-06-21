# benchmark-redis
> For benchmark redis

1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

参见: `redis_test.go`
测试结果:
begin
goos: windows
goarch: amd64
pkg: benchmark-redis
cpu: AMD Ryzen 7 5800H with Radeon Graphics         
BenchmarkRedisKv_10b
BenchmarkRedisKv_10b-16     	   34886	     35178 ns/op
BenchmarkRedisKv_20b
BenchmarkRedisKv_20b-16     	   34065	     34383 ns/op
BenchmarkRedisKv_50b
BenchmarkRedisKv_50b-16     	   32989	     35281 ns/op
BenchmarkRedisKv_100b
BenchmarkRedisKv_100b-16    	   34002	     35685 ns/op
BenchmarkRedisKv_200b
BenchmarkRedisKv_200b-16    	   32908	     35705 ns/op
BenchmarkRedisKv_1k
BenchmarkRedisKv_1k-16      	   33211	     36843 ns/op
BenchmarkRedisKv_5k
BenchmarkRedisKv_5k-16      	   28593	     40895 ns/op
PASS
end

单个key的get和set性能差距不大;


2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
1w key:
         10                 20              50              100             200         1000            5000
之前      1672408           1119960         1612336         1120912          1122472     1121648        1121648
之后      1970248           2051200         2371200         2931200          4051728     12051936       52991016
avg      29.78             93.124           75.89           181.03          292.93      1093.03         5186.94
