package databaseload

import (
	"github.com/Cynthia/goapi/pkg/database"
	logger "github.com/Cynthia/goapi/pkg/log"
	"github.com/Cynthia/goapi/pkg/model"
)

func InitDB(){
	database.Connect(logger.NewGormLogger())
	database.DB.AutoMigrate(&model.User{})
	
}