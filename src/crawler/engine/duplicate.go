package engine

import (
	"github.com/astaxie/beego/config"
	"github.com/go-redis/redis"
	"hash/crc32"
)

var client *redis.Client

var crawlerName string

func init() {
	appConfig, _ := config.NewConfig("ini", "/Users/prince/Desktop/prince/goServer/conf/logServer.conf")

	host := appConfig.String("RedisHost")
	Db, _ := appConfig.Int("RedisDB")
	crawlerName = appConfig.String("crawlerName")

	client = redis.NewClient(&redis.Options{
		Addr: host,
		DB:   Db,
	})
}

// URL去重（要修改成redis）
func isDuplicate(r Request) bool {
	if r.IsDuplicate == 1 {
		url := hashString(r.Url)

		ret := isExist(url)
		if ret == 1 {
			return true
		}
		addFilter(url)
		return false
	}

	return false
}

func addFilter(hashUrl int64) {

	redisKey := "crawler:" + crawlerName

	client.SetBit(redisKey, hashUrl, 1)
}

func isExist(hashUrl int64) int64 {
	redisKey := "crawler:" + crawlerName
	cmd := client.GetBit(redisKey, hashUrl)
	val, _ := cmd.Result()
	return val
}

func hashString(str string) int64 {
	v := int64(crc32.ChecksumIEEE([]byte(str)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
