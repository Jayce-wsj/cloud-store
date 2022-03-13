package handler

// Result Api返回结果
type Result struct {
	Status int32       `json:"status"`         //接口返回状态码
	Msg    string      `json:"msg"`            //接口返回信息
	Data   interface{} `json:"data,omitempty"` //接口返回数据
}