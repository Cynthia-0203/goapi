package validators

import (
	"github.com/Cynthia/goapi/pkg/captcha"
	"github.com/Cynthia/goapi/pkg/verifycode"
)

func ValidateCaptcha(captchaID,captchaAnswer string,errs map[string][]string)map[string][]string{
	if ok:=captcha.NewCaptcha().VerifyCaptcha(captchaID,captchaAnswer);!ok{
		
		errs["captcha_answer"]=append(errs["captcha_answer"], "captcha is wrong")
	}
	return errs
}

func ValidatePasswordConfirm(password,passwordConfirm string,errs map[string][]string)map[string][]string{
	if password!=passwordConfirm{
		errs["password_confirm"]=append(errs["password_confirm"], "passworm not match")
	}
	return errs
}

func ValidateVerifyCode(key,answer string,errs map[string][]string)map[string][]string{
	if ok:=verifycode.NewVerifyCode().CheckAnswer(key,answer);!ok{
		errs["verify_code"]=append(errs["verify_code"], "verifycode is wrong")
	}
	return errs
}