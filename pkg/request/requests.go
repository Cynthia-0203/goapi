package requests

import (
	

	"github.com/Cynthia/goapi/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{},*gin.Context)map[string][]string

func Validate(c *gin.Context,obj interface{},handler ValidatorFunc)bool{
	
	if err:=c.ShouldBind(obj);err!=nil{
		
		response.BadRequest(c,err,"failed to parse request...")
		return false
	}
	
	errs:=handler(obj,c)
	
	if len(errs)>0{
		
		response.ValidationError(c,errs)
		return false
	}
	return true
}

func validate(data interface{},rules govalidator.MapData,message govalidator.MapData)map[string][]string{
	opts:=govalidator.Options{
		Data:data,
		Rules:rules,
		TagIdentifier:"valid",
		Messages:message,
	}
	
	return govalidator.New(opts).ValidateStruct()
}

