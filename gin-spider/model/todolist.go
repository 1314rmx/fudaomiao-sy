package model

type Todolist struct {
	Id     string `gorm:"column:id;type:varchar(255);comment:Date.now()可以获得当前的时间戳;" json:"id"` // Date.now()可以获得当前的时间戳
	Title  string `gorm:"column:title;type:varchar(255);" json:"title"`
	Status int32  `gorm:"column:status;type:tinyint;default:0;" json:"status"`
	StuId  string `gorm:"column:stuId;type:varchar(255);" json:"stuId"`
	School string `gorm:"column:school;type:varchar(30);" json:"school"`
}
