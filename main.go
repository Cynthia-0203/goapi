package main

import (
	"fmt"
	
	"github.com/Cynthia/goapi/config"

	"github.com/Cynthia/goapi/pkg/database/databaseload"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/redis"
	"github.com/Cynthia/goapi/pkg/route"
	"github.com/Cynthia/goapi/pkg/validators"
	"github.com/gin-gonic/gin"

	
	"github.com/spf13/viper"
)

func main(){
	
	config.InitConfig()
	
	
	
	r:=gin.New()
	route.InitRoute(r)
	
	logger.InitLogger()
	databaseload.InitDB()
	redis.InitRedis()
	validators.InitRules()
	gin.SetMode(gin.ReleaseMode)
	
     
	
	err:=r.Run(":"+viper.GetString("app.port"))

	
	if err!=nil{
		fmt.Printf("fail to start web service...")
	}

	
}