package captcha

import (
	
	"sync"

	"github.com/Cynthia/goapi/pkg/app"
	"github.com/Cynthia/goapi/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
)

type Captcha struct{
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once
var internalCaptcha *Captcha

func NewCaptcha()*Captcha{
	once.Do(func(){
		internalCaptcha=&Captcha{}
		store:=RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix: viper.GetString("app.name")+":captcha",
		}

		driver:=base64Captcha.NewDriverDigit(
			viper.GetInt("captcha.height"),
			viper.GetInt("captcha.width"),
			viper.GetInt("captcha.length"),
			
			viper.GetFloat64("captcha.maxskew"),
			viper.GetInt("captcha.dotcount"),
		)

		internalCaptcha.Base64Captcha=base64Captcha.NewCaptcha(driver,&store)
	
	})

	return internalCaptcha
} 

func (c *Captcha)GenerateCaptcha()(id string,b64s string,answer string,err error){
	return c.Base64Captcha.Generate()
}

func(c *Captcha)VerifyCaptcha(id,answer string)bool{
	if !app.IsProduction()&&id==viper.GetString("captcha.testing_key"){
	
		return true
	}
	
	return c.Base64Captcha.Verify(id,answer,false)
}
