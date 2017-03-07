package db

import (
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"menteslibres.net/gosexy/redis"
	"strconv"
)

var instance2 *redis.Client
var once2 sync.Once

func RedisInstance() *redis.Client {
	once2.Do(func() {
		instance2 = redis.New()
		if err := instance2.Connect("docker", 6379); err != nil {
			panic(err)
		}
	})
	return instance2
}

func GetLastErshoufangProcessId() int64 {
	client := RedisInstance()
	str, _ := client.Get("process:ershoufang")
	res64, _ := strconv.ParseInt(str, 10, 64)
	return res64
}

func SetLastErshoufangProcessId(val int64) {
	client := RedisInstance()
	client.Set("process:ershoufang", val)
}

