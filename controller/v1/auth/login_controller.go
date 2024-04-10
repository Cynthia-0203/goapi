package auth

import (
	

	v1 "github.com/Cynthia/goapi/controller/v1"

	"github.com/Cynthia/goapi/pkg/auth"
	"github.com/Cynthia/goapi/pkg/jwt"
	requests "github.com/Cynthia/goapi/pkg/request"
	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
)


type LoginController struct{
	v1.BaseAPIController
}

func(lc *LoginController)LoginByPhone(c *gin.Context){
	request:=requests.LoginByPhoneRequest{}
	if ok:=requests.Validate(c,&request,requests.LoginByPhone);!ok{
		return
	}

	user,err:=auth.LoginByPhone(request.Phone)
	if err!=nil{
		response.Error(c,err,"account dose not exist")
	}else {
		token:=jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.JSON(c,gin.H{
			"token":token,
		})
	}

}

func(lc *LoginController)LoginByPassword(c *gin.Context){
	request:=requests.LoginByPasswordRequest{}
	
	
	if ok:=requests.Validate(c,&request,requests.LoginByPassword);!ok{
		return
	}

	user,err:=auth.Attempt(request.LoginID,request.Password)
	if err!=nil{
		response.Unauthorized(c,"account dose not exist or password is wrong")
	}else {
		token:=jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.JSON(c,gin.H{
			"token":token,
		})
	}

}

func(lc *LoginController)RefreshToken(c *gin.Context){
	token,err:=jwt.NewJWT().RefreshToken(c)
	if err!=nil{
		response.Error(c,err,"failed to refresh token")
	}else{
		response.JSON(c,gin.H{
			"token":token,
		})
	}
}