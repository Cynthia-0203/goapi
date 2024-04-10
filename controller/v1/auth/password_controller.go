package auth

import (
	"fmt"

	v1 "github.com/Cynthia/goapi/controller/v1"
	"github.com/Cynthia/goapi/pkg/model/user"
	requests "github.com/Cynthia/goapi/pkg/request"
	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
)

type PasswordController struct{
	v1.BaseAPIController
}

func (pc *PasswordController)ResetByPhone(c *gin.Context){
	request:=requests.ResetByPhoneRequest{}
	
	if ok:=requests.Validate(c,&request,requests.ResetByPhone);!ok{
		return
	}

	userModel:=user.GetByPhone(request.Phone)
	if userModel.ID==0{
		response.Abort404(c)
	}else{
		userModel.Password=request.Password
		userModel.Save()
		response.Success(c)
	}
}

func(pc *PasswordController)ResetByEmail(c *gin.Context){
	// 1. 验证表单
    request := requests.ResetByEmailRequest{}
    if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
        return
    }

	fmt.Println(request.Email)
    userModel := user.GetByMulti(request.Email)
    if userModel.ID == 0 {
        response.Abort404(c)
    } else {
        userModel.Password = request.Password
        userModel.Save()
        response.Success(c)
    }
}