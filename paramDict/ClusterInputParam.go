package paramDict

type ClusterInputParam struct {
	Options *ClusterInputSource
	Key    string `json:"Key"`
	Action    string `json:"action"`
	Field    string `json:"field"`
	Val    interface{} `json:"val"`
}

//func NewClusterInputParam(options *ClusterInputSource,hashParam *HashHttpInputParam) *ClusterInputParam {
//	return &ClusterInputParam{
//		Options: options,
//		Key:     hashParam.Key,
//		Field:   hashParam.Field,
//		Val:     hashParam.Val,
//	}
//}
