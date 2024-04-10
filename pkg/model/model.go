package model

import "github.com/spf13/cast"

// import "time"

type BaseModel struct{
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	// CreateTime time.Time `gorm:"column:createTime;index;" json:"createTime,omitempty"`
	// UpdateTime time.Time `gorm:"column:updateTime;index;" json:"updateTime,omitempty"`
}

func(a BaseModel)GetStringID()string{
	return cast.ToString(a.ID)
}