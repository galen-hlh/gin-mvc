package response

import (
	"encoding/json"
)

// 全局统一的响应体
type HttpResponse struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func (h *HttpResponse) ToMap() map[string]interface{} {
	var m map[string]interface{}

	j, err := json.Marshal(h)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil
	}

	return m
}
