package model

import (
	"github.com/Cynthia/goapi/pkg/hash"
	"gorm.io/gorm"
)

func(userModel *User)BeforeSave(tx *gorm.DB)(err error){
	if !hash.BcryptIsHashed(userModel.Password){
		userModel.Password=hash.BcryptHash(userModel.Password)
	}
	return
}