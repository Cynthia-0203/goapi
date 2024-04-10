package requests

import (
	"github.com/Cynthia/goapi/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct{
	Phone string`json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{},c *gin.Context)map[string][]string{
	rules:=govalidator.MapData{
		"phone":[]string{
			"required","digits:11",
		},
	}
	message:=govalidator.MapData{
		"phone":[]string{
			"required:must input phone number...",
			"digits:the len of number must be 11",
		},
	}

	return validate(data,rules,message)
}

type SignupEmailExistRequest struct{
	Email string`json:"email,omitempty" valid:"email"`
}

func ValidateSignupEmailExist(data interface{},c *gin.Context)map[string][]string{
	
	rules:=govalidator.MapData{
		"email":[]string{
			"required","min:4","max:30","email",
		},
	}
	message:=govalidator.MapData{
		"email":[]string{
			"required:must input email...",
			"min:the len of email must greater than 4...",
			"max:the len of email must less than 30...",
			"email:the format of email is not right...",
		},
	}
	
	
	return validate(data,rules,message)
	
}

type SignupUsingPhoneRequest struct{
	Phone string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Name string `json:"name" valid:"name"`
	Password string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

func SignupUsingPhone(data interface{},c *gin.Context)map[string][]string{
	rules:=govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
        "name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
        "password":         []string{"required", "min:6"},
        "password_confirm": []string{"required"},
        "verify_code":      []string{"required", "digits:6"},
	}

	messages:=govalidator.MapData{
		"phone":[]string{
			"required:must provide phone number",
			"digits:phone number must be 11 digits",
		},
		"name":[]string{
			"required:must provide name",
			"alpha_num:the format of name is wrong,must be digits or english",
			"between:the len of name between 3 and 20",
		},
		"password": []string{
            "required:must provide password",
            "min:the len of password must greater than 6",
        },
        "password_confirm": []string{
            "required:must provide password_confirm",
        },
        "verify_code": []string{
            "required:must provide verify_code",
            "digits:verify_code must be 6 digits",
        },
	}

	errs:=validate(data,rules,messages)

	_data:=data.(*SignupUsingPhoneRequest)
	errs=validators.ValidatePasswordConfirm(_data.Password,_data.PasswordConfirm,errs)
	errs=validators.ValidateVerifyCode(_data.Phone,_data.VerifyCode,errs)
	return errs
}

type SignupUsingEmailRequest struct{
	Email           string `json:"email,omitempty" valid:"email"`
    VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
    Name            string `valid:"name" json:"name"`
    Password        string `valid:"password" json:"password,omitempty"`
    PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}

func SignupUsingEmail(data interface{},c *gin.Context)map[string][]string{
	rules := govalidator.MapData{
        "email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
        "name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
        "password":         []string{"required", "min:6"},
        "password_confirm": []string{"required"},
        "verify_code":      []string{"required", "digits:6"},
    }

	messages:=govalidator.MapData{
		"email":[]string{
			"required:must provide email",
			"min:the len of email must greater than 4",
			"max:the len of email must less than 30",
			"email:the format of email is wrong",
			"not_exists:email is occupied",
		},
		"name":[]string{
			"required:must provide name",
			"alpha_num:the format of name is wrong,must be digits or english",
			"between:the len of name between 3 and 20",
		},
		"password": []string{
            "required:must provide password",
            "min:the len of password must greater than 6",
        },
        "password_confirm": []string{
            "required:must provide password_confirm",
        },
        "verify_code": []string{
            "required:must provide verify_code",
            "digits:verify_code must be 6 digits",
        },
	}
	errs:=validate(data,rules,messages)
	_data:=data.(*SignupUsingEmailRequest)
	errs=validators.ValidatePasswordConfirm(_data.Password,_data.PasswordConfirm,errs)
	errs=validators.ValidateVerifyCode(_data.Email,_data.VerifyCode,errs)
	return errs
}