package auth

import (
	"fmt"
	"net/http"

	v1 "github.com/Cynthia/goapi/controller/v1"
	"github.com/Cynthia/goapi/pkg/database"
	"github.com/Cynthia/goapi/pkg/jwt"
	"github.com/Cynthia/goapi/pkg/model"
	"github.com/Cynthia/goapi/pkg/model/user"
	requests "github.com/Cynthia/goapi/pkg/request"
	"github.com/Cynthia/goapi/pkg/response"

	"github.com/gin-gonic/gin"
)

type SignupController struct{
	v1.BaseAPIController
}

func (sc *SignupController)IsPhoneExist(c *gin.Context){
	
	
	request:=requests.SignupPhoneExistRequest{}

	if ok:=requests.Validate(c,&request,requests.ValidateSignupPhoneExist);!ok{
		return
	}
	
	

	response.JSON(c,gin.H{
		"exist":user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController)IsEmailExist(c *gin.Context){
	request:=requests.SignupEmailExistRequest{}

	err:=bind(&request,c)
	if err!=nil{
		fmt.Println(err)
		return
	}

	errs:=requests.ValidateSignupEmailExist(&request,c)
	if len(errs)>0{
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"errors":errs,
		})
		return
	}
	
	
	response.JSON(c,gin.H{
		"exist":user.IsEmailExist(request.Email),
	})
}

func bind(data interface{},c *gin.Context)error{
	if err:=c.ShouldBindJSON(&data);err!=nil{
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"error":err.Error(),
		})
		return err
	}
	return nil
}

func(sc *SignupController)SignupUsingPhone(c *gin.Context){
	request:=requests.SignupUsingPhoneRequest{}
	if ok:=requests.Validate(c,&request,requests.SignupUsingPhone);!ok{
		return
	}

	_user:=model.User{
		Name: request.Name,
		Phone: request.Phone,
		Password: request.Password,
	}

	database.DB.Create(&_user)
	if _user.ID>0{
		token:=jwt.NewJWT().IssueToken(_user.GetStringID(),_user.Name)
		response.CreatedJSON(c,gin.H{
			"token":token,
			"data":_user,
		})
	}else{
		response.Abort500(c,"failed to create user...")
	}
	
}

func(sc *SignupController)SignupUsingEmail(c *gin.Context){
	request:=requests.SignupUsingEmailRequest{}
	
	if ok:=requests.Validate(c,&request,requests.SignupUsingEmail);!ok{
		return
	}

	user:=model.User{
		Name: request.Name,
		Phone: request.Email,
		Password: request.Password,
	}

	database.DB.Create(&user)

	if user.ID>0{
		token:=jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.CreatedJSON(c,gin.H{
			"data":user,
			"token":token,
		})
	}else{
		response.Abort500(c,"failed to create user...")
	}
	
}