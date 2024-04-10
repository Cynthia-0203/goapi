package user

import (
	
	"github.com/Cynthia/goapi/pkg/database"
	"github.com/Cynthia/goapi/pkg/model"
)

func IsEmailExist(email string)bool{
	var count int64
	database.DB.Model(model.User{}).Where("email = ?",email).Count(&count)
	return count>0
}

func IsPhoneExist(phone string)bool{
	var count int64
	database.DB.Model(model.User{}).Where("phone = ?",phone).Count(&count)
	return count>0
}

func GetByPhone(phone string)(userModel model.User){
	database.DB.Where("phone=?",phone).First(&userModel)
	return
}

func GetByMulti(loginID string)(userModel model.User){
	database.DB.Where("phone=?",loginID).Or("email=?",loginID).Or("name=?",loginID).First(&userModel)
	return
}

func Get(idstr string)(userModel model.User){
	database.DB.Where("id=?",idstr).First(&userModel)
	return
}

// func GetByEmail(email string)(userModel model.User){
// 	database.DB.Where("email=?",email).First(&userModel)
// 	return
// }