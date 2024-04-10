package middlewares

import (
	"fmt"
	"github.com/Cynthia/goapi/pkg/jwt"
	"github.com/Cynthia/goapi/pkg/model/user"
	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthJWT()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		claims,err:=jwt.NewJWT().ParserToken(ctx)

		if err!=nil{
			response.Unauthorized(ctx,fmt.Sprintf("view interface documentation of %v",viper.GetString("app.name")))
			return
		}

		userModel:=user.Get(claims.Id)
		if userModel.ID==0{
			response.Unauthorized(ctx,"failed to find user")
			return
		}
		ctx.Set("current_user_id",userModel.GetStringID())
		ctx.Set("current_user_name",userModel.Name)
		ctx.Set("current_user",userModel)
		ctx.Next()
	}
}