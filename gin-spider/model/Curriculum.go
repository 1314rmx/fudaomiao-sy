package model

type Curriculum struct {
	Xsxx struct {
		Kcms int `json:"KCMS"` //课程名
	} `json:"xsxx"`
	XqjmcMap struct {
		Num1 string `json:"1"`
		Num2 string `json:"2"`
		Num3 string `json:"3"`
		Num4 string `json:"4"`
		Num5 string `json:"5"`
		Num6 string `json:"6"`
		Num7 string `json:"7"`
	} `json:"xqjmcMap"` //星期几
	KbList []struct {
		Cdlbmc    string `json:"cdlbmc"` //教室类型
		Cdmc      string `json:"cdmc"`   //上课教室
		Jc        string `json:"jc"`     //1-2节
		Jxbmc     string `json:"jxbmc"`  //课程名
		Jxbsftkbj string `json:"jxbsftkbj"`
		Jxbzc     string `json:"jxbzc"`  //上课的班级
		Kcxszc    string `json:"kcxszc"` //授课的时间
		Xf        string `json:"xf"`     //学分
		Xm        string `json:"xm"`     //上课老师
		Xqjmc     string `json:"xqjmc"`  //上课星期几
		Xqmc      string `json:"xqmc"`   //校区
		Zcd       string `json:"zcd"`    //上课的周数
		Zxs       string `json:"zxs"`    //总学时
	}
}
