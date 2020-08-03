package domain

//查询 es查询
type EsData struct {
	Total int64       `json:"total"`
	Items interface{} `json:"items"`
}

type EsResult struct {
	Product    []string `json:"product"`
	ProductNum int64    `json:"productNum"`
	LiveNum    int64    `json:"liveNum"`
	Live       []string `json:"live"`
	MemberNum  int64    `json:"memberNum"`
	Member     []string `json:"member"`
}
