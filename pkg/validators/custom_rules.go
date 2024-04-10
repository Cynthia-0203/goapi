package validators

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Cynthia/goapi/pkg/database"
	"github.com/thedevsaddam/govalidator"
)


func InitRules(){
	govalidator.AddCustomRule("not_exists",func(field, rule, message string, value interface{}) error {
		fmt.Println(rule)
		rng:=strings.Split(strings.TrimPrefix(rule,"not_exists:"),",")
		
		tableName:=rng[0]
		fmt.Println(tableName)
		dbFiled:=rng[1]
		var exceptID string
		if len(rng)>2{
			exceptID=rng[2]
		}

		requestValue:=value.(string)
		query:=database.DB.Table(tableName).Where(dbFiled+"=?",requestValue)
		if len(exceptID)>0{
			query.Where("id!=?",exceptID)
		}
		var count int64
		query.Count(&count)
		if count!=0{
			if message!=""{
				return errors.New(message)
			}
			return fmt.Errorf("%v is occupied",requestValue)
		}
		return nil
	})
}