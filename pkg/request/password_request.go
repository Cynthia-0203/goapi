package requests

import (
	"fmt"

	"github.com/Cynthia/goapi/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct{
	Phone      string `json:"phone,omitempty" valid:"phone"`
    VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
    Password   string `json:"password,omitempty" valid:"password"`
}

func ResetByPhone(data interface{},c *gin.Context)map[string][]string{
	rules := govalidator.MapData{
        "phone":       []string{"required", "digits:11"},
        "verify_code": []string{"required", "digits:6"},
        "password":    []string{"required", "min:6"},
    }
    messages := govalidator.MapData{
        "phone": []string{
            "required:must provide phone",
            "digits:phone must be 11 digits",
        },
        "verify_code": []string{
            "required:must provide verify_code",
            "digits:verify_code must be 6 digits",
        },
        "password": []string{
            "required:must provide password",
            "min:the len of password must greater than 6",
        },
    }

    errs := validate(data, rules, messages)
	fmt.Println(errs)
    
    _data := data.(*ResetByPhoneRequest)
    errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	fmt.Println(errs)
    return errs
}

type ResetByEmailRequest struct{
	Email      string `json:"email,omitempty" valid:"email"`
    VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
    Password   string `valid:"password" json:"password,omitempty"`
}

func ResetByEmail(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "email":       []string{"required", "min:4", "max:30", "email"},
        "verify_code": []string{"required", "digits:6"},
        "password":    []string{"required", "min:6"},
    }
    messages := govalidator.MapData{
        "email": []string{
            "required:must provide mail",
            "min:the len of email must greater than 4",
            "max:the len of email must less than 30",
            "email:the format of email is wrong",
        },
        "verify_code": []string{
            "required:must provide verify_code",
            "digits:verify_code must be 6 digits",
        },
        "password": []string{
            "required:must provide password",
            "min:the len of password must greater than 6",
        },
    }

    errs := validate(data, rules, messages)

    // 检查验证码
    _data := data.(*ResetByEmailRequest)
    errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

    return errs
}