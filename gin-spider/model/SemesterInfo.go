package model

type SemesterInfo struct {
	NjdmID string `json:"njdm_id"` //年级
	Xz     string `json:"xz"`      //几年制
	Xm     string `json:"xm"`      //姓名
	Xh     string `json:"xh"`      //学号
	Bhid   string `json:"bh_id"`   //专业班级
	Zsjgid string `json:"zsjg_id"` //学院
}
