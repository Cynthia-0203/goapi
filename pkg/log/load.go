package logger

import "github.com/spf13/viper"

func InitLogger(){
	filename:=viper.GetString("log.filename")
	maxSize,maxAge,maxBackup:=viper.GetInt("log.max_size"),viper.GetInt("log.max_age"),viper.GetInt("log.max_backup")
	compress:=viper.GetBool("log.compress")
	level:=viper.GetString("log.level")
	log_type:=viper.GetString("log.type")
	logger(filename,maxSize,maxAge,maxBackup,compress,log_type,level)
}