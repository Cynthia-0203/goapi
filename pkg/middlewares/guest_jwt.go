package middlewares

import (
	"github.com/Cynthia/goapi/pkg/jwt"
	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
)

func GuestJWT()gin.HandlerFunc{
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization"))>0{
			_,err:=jwt.NewJWT().ParserToken(c)

			if err==nil{
				response.Unauthorized(c,"please use tourist status to login")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}