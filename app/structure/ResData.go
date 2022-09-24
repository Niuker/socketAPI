package structure

type ResData struct {
	Code      int         `json:"code"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
	Timestamp int         `json:"timestamp"`
	Reqid     string      `json:"reqid"`
}
