package app

import "github.com/spf13/viper"

func IsLocal()bool{
	return viper.GetString("app.env")=="local"
}

func IsProduction()bool{
	return viper.GetString("app.env")=="produciton"
}

func IsTesting()bool{
	return viper.GetString("app.env")=="Testing"
}