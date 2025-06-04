package request

type WsRequest struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

//	type Online struct {
//		UserId int64 `json:"userId"`
//	}
type OnlineRequest struct {
	Cmd string `json:"cmd"`
	//Data Online `json:"data"`
}

type Send struct {
	//FormUserId int64  `json:"formUserId"`
	ToUserId int64  `json:"toUserId"`
	Message  string `json:"message"`
}
type SendRequest struct {
	Cmd  string `json:"cmd"`
	Data Send   `json:"data"`
}

type AllSend struct {
	Message string `json:"message"`
}
type AllSendRequest struct {
	Cmd  string  `json:"cmd"`
	Data AllSend `json:"data"`
}
