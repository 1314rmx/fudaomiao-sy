package model

type SemesterInfo struct {
	NjdmID string `json:"njdm_id"` //年级
	Xz     string `json:"xz"`      //几年制
	Xm     string `json:"xm"`      //姓名
	Xh     string `json:"xh"`      //学号
}
