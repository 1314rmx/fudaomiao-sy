package model

type ClassRoom struct {
	Items []struct {
		Cdlbmc string `json:"cdlbmc"` //教室类型
		Cdmc   string `json:"cdmc"`   //教室位置
		Jgmc   string `json:"jgmc"`   //教室属于那个学院
		Jxlmc  string `json:"jxlmc"`  //教学楼
		Zws    string `json:"zws"`    //座位数
	} `json:"items"`
}
