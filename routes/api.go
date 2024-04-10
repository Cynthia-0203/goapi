package routes

import (
	"github.com/Cynthia/goapi/controller/v1/auth"
	"github.com/Cynthia/goapi/pkg/middlewares"
	"github.com/gin-gonic/gin"
)
func RegisterAPIRoutes(r *gin.Engine){
	v1:=r.Group("/v1")

	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup:=v1.Group("/auth")
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc:=new(auth.SignupController)
			authGroup.POST("/signup/phone/exist",middlewares.GuestJWT(),suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist",middlewares.GuestJWT(),suc.IsEmailExist)
			authGroup.POST("/signup/using-phone",middlewares.GuestJWT(),suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email",middlewares.GuestJWT(),suc.SignupUsingEmail)
			vcc:=new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha",middlewares.LimitPerRoute("20-H"),vcc.ShowCaptcha)

			authGroup.POST("/verify-codes/phone",middlewares.LimitPerRoute("20-H"),vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email",middlewares.LimitPerRoute("20-H"),vcc.SendUsingEmail)
			lc:=new(auth.LoginController)
			authGroup.POST("/login/using-phone",middlewares.GuestJWT(),lc.LoginByPhone)
			authGroup.POST("/login/using-password",middlewares.GuestJWT(),lc.LoginByPassword)
			authGroup.POST("/login/refresh-token",middlewares.AuthJWT(),lc.RefreshToken)

			pwc:=new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone",middlewares.GuestJWT(),pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email",middlewares.GuestJWT(),pwc.ResetByEmail)
		}
	}
}