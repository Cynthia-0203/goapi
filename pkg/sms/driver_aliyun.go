package sms

import (
	"encoding/json"

	logger "github.com/Cynthia/goapi/pkg/log"
	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"github.com/spf13/viper"
)


type Aliyun struct{}

func(s *Aliyun)Send(phone string,message Message,config map[string]string)bool{
	smsClient:=aliyunsmsclient.New("http://dysmsapi.aliyuncs.com/")

	templateParam,err:=json.Marshal(message.Data)
	if err!=nil{
		logger.ErrorString("message[aliyun]","failed to parse data",err.Error())
	}

	logger.DebugJSON("message[aliyun]","configuration information",config)

	result,err:=smsClient.Execute(
		viper.GetString("aliyun.key_id"),
		viper.GetString("aliyun.key_secret"),
		phone,
		viper.GetString("aliyun.sign_name"),
		message.Template,
		string(templateParam),
	)
	logger.DebugJSON("message[aliyun]","request content",smsClient.Request)
	logger.DebugJSON("message[aliyun]","interface response",result)

	if err!=nil{
		logger.ErrorString("message[aliyun]","failed to send message...",err.Error())
	}
	resultJSON,err:=json.Marshal(result)
	if  err!=nil {
		logger.ErrorString("message[aliyun]","failed to parse json...",err.Error())
		return false
	}
	if result.IsSuccessful(){
		logger.DebugString("message[aliyun]","success to send message...","")
		return true
	}else{
		logger.ErrorString("message[aliyun]","server error",string(resultJSON))
		return false
	}
}