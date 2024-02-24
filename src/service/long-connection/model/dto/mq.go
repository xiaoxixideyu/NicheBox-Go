package dto

type PushToUserMessage struct {
	Uid  int64  `json:"uid"`
	Data []byte `json:"data"`
}

type PushToDeviceMessage struct {
	Uid       int64  `json:"uid"`
	UserAgent string `json:"user_agent"`
	Data      []byte `json:"data"`
}
