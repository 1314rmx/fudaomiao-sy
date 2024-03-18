package model

type Stuscore struct {
	Items []struct {
		Bj     string `json:"bj"`     //班级
		Cj     string `json:"cj"`     //成绩
		Jd     string `json:"jd"`     //绩点
		Jsxm   string `json:"jsxm"`   //老师
		Kcxzmc string `json:"kcxzmc"` //类型
		Sfxwkc string `json:"sfxwkc"` //是否学位课程
		Xf     string `json:"xf"`     //学分
		Xfjd   string `json:"xfjd"`   //学分绩点
		Xnmmc  string `json:"xnmmc"`  //2023-2024学年
		Xqmmc  string `json:"xqmmc"`  //学期数
		Kcmc   string `json:"kcmc"`   //课程名
		Khfsmc string `json:"khfsmc"` //考试方式
		Xm     string `json:"xm"`     //姓名
		Xh     string `json:"xh"`     //学号
		Kkbmmc string `json:"kkbmmc"` //开课学院
	} `json:"items"`
}
