package model

import (
	"github.com/Cynthia/goapi/pkg/database"
	"github.com/Cynthia/goapi/pkg/hash"
)



type User struct{
	BaseModel
	Name string`json:"name,omitempty"`
	Email string`json:"-"`
	Phone string`json:"-"`
	Password string`json:"-"`
}

func (userModel *User)ComparePassword(_password string)bool{
	return hash.BcryptCheck(_password,userModel.Password)
}

func(userModel *User)Save()(rowAffected int64){
	result:=database.DB.Save(&userModel)
	return result.RowsAffected
}

