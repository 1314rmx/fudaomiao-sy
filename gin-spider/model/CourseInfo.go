package model

type CourseInfo struct {
	Items []struct {
		JghID string `json:"jgh_id"` //id
		JxbID string `json:"jxb_id"` //id
		KchID string `json:"kch_id"` //id
		Kcmc  string `json:"kcmc"`   //kcmc
		Tjzt  string `json:"tjzt"`   //提交状态
	} `json:"items"`
}
