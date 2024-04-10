package limiter

import (
	"strings"

	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(c *gin.Context)string{
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context)string{
	return routeToKeyString(c.FullPath()+c.ClientIP())
}

func routeToKeyString(routeName string)string{
	routeName=strings.ReplaceAll(routeName,"/","-")
	routeName=strings.ReplaceAll(routeName,":","_")
	return routeName
}

func CheckRate(c *gin.Context,key string,formatted string)(limiterlib.Context,error){
	var context limiterlib.Context
	rate,err:=limiterlib.NewRateFromFormatted(formatted)
	if err!=nil{
		logger.LogIf(err)
		return context,err
	}

	store,err:=sredis.NewStoreWithOptions(redis.Redis.Client,limiterlib.StoreOptions{
		Prefix: viper.GetString("app.name")+":limiter",
	})

	if err!=nil{
		logger.LogIf(err)
		return context,err
	}

	limiterObj:=limiterlib.New(store,rate)

	if c.GetBool("limiter-once"){
		return limiterObj.Peek(c,key)
	}else{
		c.Set("limiter-once",true)
		return limiterObj.Get(c,key)
	}

}