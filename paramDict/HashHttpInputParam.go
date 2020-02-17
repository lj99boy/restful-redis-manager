package paramDict

type HashHttpInputParam struct {
	Key    string `json:"Key"`
	Action    string `json:"action"`
	Field    string `json:"field"`
	Val    interface{} `json:"val"`
}
