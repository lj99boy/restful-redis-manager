package paramDict

type KeyHttpInputParam struct {
	Key string
	Action string `json:"action"`
	Val    interface{} `json:"val"`
}

