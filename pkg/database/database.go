package database

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(logger gormlogger.Interface){
	var err error
	DB,err=gorm.Open(mysql.New(mysql.Config{
		DSN:fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
        viper.GetString("mysql.password"),
        viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
        viper.GetString("mysql.database"),
        viper.GetString("mysql.chareset")),
	}),&gorm.Config{
		Logger: logger,
	})
	if err!=nil{
		fmt.Println(err)
	}

	SQLDB,err=DB.DB()
	if err!=nil{
		fmt.Println("failed to create sqldb...")
	}

	SQLDB.SetConnMaxIdleTime(viper.GetDuration("mysql.max_life_seconds"))
	SQLDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_connection"))
	SQLDB.SetMaxOpenConns(viper.GetInt("mysql.max_open_connection"))
}

