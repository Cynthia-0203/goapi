package redis

import (
	"context"
	"sync"
	"time"

	logger "github.com/Cynthia/goapi/pkg/log"
	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct{
	Client *redis.Client
	Context context.Context //all service need context
}

var once sync.Once
var Redis *RedisClient

func ConnectRedis(address string,db int){
	
	once.Do(func(){
		Redis=NewClient(address,db)
	})
}

func NewClient(address string,db int)*RedisClient{
	rds:=&RedisClient{}
	rds.Context=context.Background()
	rds.Client=redis.NewClient(&redis.Options{
		Addr: address,
		// Username: username,
		// Password: password,
		DB: db,
	})
	_,err:=rds.Client.Ping(rds.Context).Result()
	logger.LogIf(err)
	return rds

}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
    if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
        logger.ErrorString("Redis", "Set", err.Error())
        return false
    }
    return true
}

// Get 获取 key 对应的 value
func (rds RedisClient) Get(key string) string {
    result, err := rds.Client.Get(rds.Context, key).Result()
    if err != nil {
        if err != redis.Nil {
            logger.ErrorString("Redis", "Get", err.Error())
        }
        return ""
    }
    return result
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds RedisClient) Has(key string) bool {
    _, err := rds.Client.Get(rds.Context, key).Result()
    if err != nil {
        if err != redis.Nil {
            logger.ErrorString("Redis", "Has", err.Error())
        }
        return false
    }
    return true
}

//Del 删除存储在 redis 里的数据，支持多个 key 传参
func (rds RedisClient) Del(keys ...string) bool {
    if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
        logger.ErrorString("Redis", "Del", err.Error())
        return false
    }
    return true
}

// FlushDB 清空当前 redis db 里的所有数据
func (rds RedisClient) FlushDB() bool {
    if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
        logger.ErrorString("Redis", "FlushDB", err.Error())
        return false
    }
    return true
}

func (rds RedisClient)Increment(parameters ...interface{})bool{
	switch len(parameters){
	case 1:
		key:=parameters[0].(string)
		if err:=rds.Client.Incr(rds.Context,key).Err();err!=nil{
			logger.ErrorString("redis","increment",err.Error())
			return false
		}
	case 2:
		key:=parameters[0].(string)
		value:=parameters[1].(int64)
		if err:=rds.Client.IncrBy(rds.Context,key,value).Err();err!=nil{
			logger.ErrorString("redis","increment",err.Error())
			return false
		}
	default:
		logger.ErrorString("redis","increment","parameters are too much...")
		return false
	}
	return true
}

func (rds RedisClient)Decrement(parameters ...interface{})bool{
	switch len(parameters){
	case 1:
		key:=parameters[0].(string)
		if err:=rds.Client.Decr(rds.Context,key).Err();err!=nil{
			logger.ErrorString("redis","decrement",err.Error())
			return false
		}
	case 2:
		key:=parameters[0].(string)
		value:=parameters[1].(int64)
		if err:=rds.Client.DecrBy(rds.Context,key,value).Err();err!=nil{
			logger.ErrorString("redis","decrement",err.Error())
			return false
		}
	default:
		logger.ErrorString("redis","decrement","parameters are too much...")
		return false
	}
	return true
}
