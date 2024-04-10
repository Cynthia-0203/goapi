package verifycode

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Cynthia/goapi/pkg"
	"github.com/Cynthia/goapi/pkg/app"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/mail"
	"github.com/Cynthia/goapi/pkg/redis"
	"github.com/Cynthia/goapi/pkg/sms"
	"github.com/spf13/viper"
)

type VerifyCode struct{
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode()*VerifyCode{
	once.Do(func() {
		internalVerifyCode=&VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix: viper.GetString("app.name")+":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode)generateVerifyCode(key string)string{
	code:=pkg.RanddomNumber(viper.GetInt("verifycode.code_length"))

	if app.IsLocal(){
		code=viper.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("captcha","generate captcha",map[string]string{key:code})
	vc.Store.Set(key,code)
	return code
}
func(vc *VerifyCode)SendSMS(phone string)bool{
	code:=vc.generateVerifyCode(phone)
	if !app.IsProduction()&&strings.HasPrefix(phone,viper.GetString("verifycode.debug_phone+prefix")){
		return true
	}
	return sms.NewSMS().Send(phone,sms.Message{
		Template: viper.GetString("aliyun.template_code"),
		Data: map[string]string{"code":code},
	})
}

func(vc *VerifyCode)SendEmail(email string)error{
	code:=vc.generateVerifyCode(email)
	
	if !app.IsProduction()&&strings.HasSuffix(email,viper.GetString("verifycode.debug_email_suffix")){
		
		return nil
	}
	content:=fmt.Sprintf("<h1>your captcha is %v</h1>",code)

	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: viper.GetString("mail.from.address"),
			Name: viper.GetString("mail.from.name"),
		},
		To: []string{email},
		Subject: "captcha of email",
		HTML: []byte(content),
	})

	return nil
}


func(vc *VerifyCode)CheckAnswer(key,answer string)bool{
	logger.DebugJSON("captcha","check answer",map[string]string{key:answer})

	if !app.IsProduction()&&(strings.HasSuffix(key,viper.GetString("verifycode.debug_email_suffix"))||strings.HasPrefix(key,viper.GetString("verifycode.debug_phone_prefix"))){
		return true
	}
	return vc.Store.Verify(key,answer,false)
}