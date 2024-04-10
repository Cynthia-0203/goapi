package auth

import (
	

	v1 "github.com/Cynthia/goapi/controller/v1"
	"github.com/Cynthia/goapi/pkg/captcha"
	logger "github.com/Cynthia/goapi/pkg/log"
	requests "github.com/Cynthia/goapi/pkg/request"
	"github.com/Cynthia/goapi/pkg/verifycode"

	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct{
	v1.BaseAPIController
}

func (vc *VerifyCodeController)ShowCaptcha(c *gin.Context){
	id,_,b64s,err:=captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(c,gin.H{
		"captcha_id":id,
		"captcha_image":b64s,
	})
}

func(vc *VerifyCodeController)SendUsingPhone(c *gin.Context){
	request:=requests.VerifyCodePhoneRequest{}

	if ok:=requests.Validate(c,&request,requests.VerifyCodePhone);!ok{
		return
	}

	if ok:=verifycode.NewVerifyCode().SendSMS(request.Phone);!ok{
		response.Abort500(c,"failed to send sms")
	}else{
		response.Success(c)
	}

   
}

func (vc *VerifyCodeController)SendUsingEmail(c *gin.Context){
	request:=requests.VerifyCodeEmailRequest{}
	if ok:=requests.Validate(c,&request,requests.VerifyCodeEmail);!ok{
		return
	}

	err:=verifycode.NewVerifyCode().SendEmail(request.Email)
	if err!=nil{
		response.Abort500(c,"failed to send captcha...")
	}else{
		response.Success(c)
	}
}