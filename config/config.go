package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init(){
    // 设置配置文件的名字
    viper.SetConfigName("config")
    // 设置配置文件的类型
    viper.SetConfigType("yaml")
    // 添加配置文件的路径，指定 config 目录下寻找
    viper.AddConfigPath(".")
    // 寻找配置文件并读取
    err := viper.ReadInConfig()
    if err != nil {
            panic(fmt.Errorf("fatal error config file: %w", err))
    } 
}
