package response

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type WsResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func WsResponseError(c *websocket.Conn, code int64, msg string) {
	res := WsResponse{
		Code: code,
		Msg:  msg,
	}
	marshal, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}
	err = c.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Println("write:", err)
		return
	}

}
func WsResponseSuccess(c *websocket.Conn, data interface{}) error {
	res := WsResponse{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
	marshal, err := json.Marshal(&res)
	if err != nil {
		panic(err)
	}
	err = c.WriteMessage(websocket.TextMessage, marshal)
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
}
