package structure

type ReqData struct {
	Event     string            `json:"event"`
	Params    map[string]string `json:"params"`
	Timestamp int               `json:"timestamp"`
	Reqid     string            `json:"reqid"`
}
