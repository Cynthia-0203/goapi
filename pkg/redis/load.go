package redis

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitRedis(){
	ConnectRedis(fmt.Sprintf("%v:%v",viper.GetString("redis.host"),viper.GetString("redis.port")),
	// viper.GetString("redis.username"),
	// viper.GetString("redis.password"),
	viper.GetInt("redis.database"))
}