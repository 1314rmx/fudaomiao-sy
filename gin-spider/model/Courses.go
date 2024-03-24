package model

type Courses struct {
	KbList []struct {
		Cdlbmc       string   `json:"cdlbmc"`       //教室类型
		Cdmc         string   `json:"cdmc"`         //上课教室
		Jc           string   `json:"jc"`           //1-2节
		Jxbmc        string   `json:"jxbmc"`        //课程名
		Khfsmc       string   `json:"khfsmc"`       //考察
		Kcxz         string   `json:"kcxz"`         //必修
		Jxbzc        string   `json:"jxbzc"`        //上课的班级
		Kcxszc       string   `json:"kcxszc"`       //授课的时间
		Xf           string   `json:"xf"`           //学分
		Xm           string   `json:"xm"`           //上课老师
		Xqjmc        string   `json:"xqjmc"`        //上课星期几
		Xqmc         string   `json:"xqmc"`         //校区
		Zcd          string   `json:"zcd"`          //上课的周数
		Weeks        []string `json:"weeks"`        //周数(手动增加)
		Zxs          string   `json:"zxs"`          //总学时
		Section      string   `json:"section"`      //开始节数(手动增加)
		SectionCount string   `json:"sectionCount"` //一共几节(手动增加)
		Week         string   `json:"week"`         //星期几(手动增加)
	}
}
