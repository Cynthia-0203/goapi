package middlewares

import (
	"net/http"

	"github.com/Cynthia/goapi/pkg/app"
	"github.com/Cynthia/goapi/pkg/limiter"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func LimitIP(limit string)gin.HandlerFunc{
	if app.IsTesting(){
		limit="1000000-H"
	}

	return func(ctx *gin.Context) {
		key:=limiter.GetKeyIP(ctx)
		if ok:=limitHandler(ctx,key,limit);!ok{
			return
		}
		ctx.Next()
	}
}

func limitHandler(c *gin.Context,key string,limit string)bool  {
	rate,err:=limiter.CheckRate(c,key,limit)
	if err!=nil{
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
    c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
    c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	if rate.Reached{
		c.AbortWithStatusJSON(http.StatusTooManyRequests,gin.H{
			"message":"requests are too many",
		})
		return false
	}
	return true
}

func LimitPerRoute(limit string) gin.HandlerFunc {
    if app.IsTesting() {
        limit = "1000000-H"
    }
    return func(c *gin.Context) {

        // 针对单个路由，增加访问次数
        c.Set("limiter-once", false)

        // 针对 IP + 路由进行限流
        key := limiter.GetKeyRouteWithIP(c)
        if ok := limitHandler(c, key, limit); !ok {
            return
        }
        c.Next()
    }
}