package model

type Evaluate struct {
	Id      int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Content string `gorm:"column:content;type:text;" json:"content"`
}
