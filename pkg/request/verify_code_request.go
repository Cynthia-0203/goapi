package requests

import (
	

	"github.com/Cynthia/goapi/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct{
	CaptchaID string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func VerifyCodePhone(data interface{},c *gin.Context)map[string][]string{
	rules := govalidator.MapData{
        "phone":          []string{"required", "digits:11"},
        "captcha_id":     []string{"required"},
        "captcha_answer": []string{"required", "digits:6"},
    }

	messages:=govalidator.MapData{
		"phone":[]string{
			"required:must provide phone number",
			"digits:the len of phone number must be 11 and all are digits",
		},
		"captcha_id":[]string{
			"required:must provide captcha id",
		},
		"captcha_answer":[]string{
			"required:must provide captcha answer",
		},
	}
	
	errs:=validate(data,rules,messages)
	
	_data:=data.(*VerifyCodePhoneRequest)
	
	errs=validators.ValidateCaptcha(_data.CaptchaID,_data.CaptchaAnswer,errs)
	return errs
}

type VerifyCodeEmailRequest struct {
    CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
    CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

    Email string `json:"email,omitempty" valid:"email"`
}


func VerifyCodeEmail(data interface{},c *gin.Context)map[string][]string{
	rules:=govalidator.MapData{
		"email":[]string{"required","min:4","max:30","email"},
		"captcha_id":[]string{"required"},
		"captcha_answer":[]string{"required","digits:6"},
	}

	messages:=govalidator.MapData{
		"email":[]string{
			"required:must provide email",
			"min:the len of email must greater than 4",
			"max:the len of email must less than 30",
			"email:the format of mail is wrong",
		},
		"captcha_id":[]string{
			"required:must provide captcha_id",
		},
		"captcha_answer":[]string{
			"required:must provide captcha_answer",
			"digits:captcha is 6 digits",
		},
	}

	errs:=validate(data,rules,messages)

	_data:=data.(*VerifyCodeEmailRequest)

	errs=validators.ValidateCaptcha(_data.CaptchaID,_data.CaptchaAnswer,errs)
	
	return errs
}

