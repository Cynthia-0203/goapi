package verifycode

import (
	"time"

	"github.com/Cynthia/goapi/pkg/app"
	"github.com/Cynthia/goapi/pkg/redis"
	"github.com/spf13/viper"
)

type RedisStore struct{
	RedisClient *redis.RedisClient
	KeyPrefix string
}

func(s *RedisStore)Set(key string,value string)bool{
	ExpireTime:=time.Minute*time.Duration(viper.GetInt64("verifycode.expire_time"))
	if app.IsLocal(){
		ExpireTime=time.Minute*time.Duration(viper.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPrefix+key,value,ExpireTime)
}

func(s *RedisStore)Get(key string,clear bool)(value string){
	key=s.KeyPrefix+key
	val:=s.RedisClient.Get(key)
	if clear{
		s.RedisClient.Del(key)
	}
	return val
}

func(s *RedisStore)Verify(key,answer string,clear bool)bool{
	v:=s.Get(key,clear)
	return v==answer
}
