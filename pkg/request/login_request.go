package requests

import (
	"github.com/Cynthia/goapi/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct{
	Phone string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "phone":       []string{"required", "digits:11"},
        "verify_code": []string{"required", "digits:6"},
    }
    messages := govalidator.MapData{
        "phone": []string{
            "required:must provide phone",
            "digits:must be 11 digits",
        },
        "verify_code": []string{
            "required:must provide verify_code",
            "digits:must be 6 digits",
        },
    }

    errs := validate(data, rules, messages)

   
    _data := data.(*LoginByPhoneRequest)
    errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

    return errs
}


type LoginByPasswordRequest struct {
    CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
    CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

    LoginID  string `valid:"login_id" json:"login_id"`
    Password string `valid:"password" json:"password,omitempty"`
}
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "login_id":       []string{"required", "min:3"},
        "password":       []string{"required", "min:6"},
        "captcha_id":     []string{"required"},
        "captcha_answer": []string{"required", "digits:6"},
    }
    messages := govalidator.MapData{
        "login_id": []string{
            "required:must provide email or phone or username",
            "min:the len of id must greater than 3",
        },
        "password": []string{
            "required:must provide password",
            "min:the len of password must greater than 6",
        },
        "captcha_id": []string{
            "required:must provide captcha_id",
        },
        "captcha_answer": []string{
            "required:must provide captcha_answer",
            "digits:must be 6 digits",
        },
    }

    errs := validate(data, rules, messages)

    // 图片验证码
    _data := data.(*LoginByPasswordRequest)
    errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

    return errs
}