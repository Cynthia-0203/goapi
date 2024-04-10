package auth

import (
	"errors"

	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/model"
	"github.com/Cynthia/goapi/pkg/model/user"
	"github.com/gin-gonic/gin"
)

func Attempt(email,password string)(model.User,error){
	userModel:=user.GetByMulti(email)
	if userModel.ID==0{
		return model.User{},errors.New("account not exists")
	}

	if !userModel.ComparePassword(password){
		return model.User{},errors.New("password is wrong")
	}

	return userModel,nil
}

func LoginByPhone(phone string)(model.User,error){
	userModel:=user.GetByPhone(phone)
	if userModel.ID==0{
		return model.User{},errors.New("this phone is unregistered")
	}

	return userModel,nil
}

func CurrentUser(c *gin.Context)model.User{
	uerModel,ok:=c.MustGet("current_user").(model.User)
	if !ok{
		logger.LogIf(errors.New("failed to get user"))
		return model.User{}
	}

	return uerModel
}

func CurrentUID(c *gin.Context)string{
	return c.GetString("current_user_id")
}