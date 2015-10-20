package myim

import (
	"encoding/json"
	"golang.org/x/net/websocket"
)

type ReqJsonRpc struct {
	Method string            `json:"method"`
	Params map[string]string `json:"params"`
}

type SenderStruct struct {
}

var (
	Sender = &SenderStruct{}
)

func (sender *SenderStruct) Error(ws *websocket.Conn, num, message string) {
	var params = make(map[string]string)
	params["id"] = num
	params["message"] = message
	var reqJson = ReqJsonRpc{
		"error",
		params,
	}
	sender.Send(ws, reqJson)
}

func (sender *SenderStruct) Notice(ws *websocket.Conn, message string) {
	var params = make(map[string]string)
	params["message"] = message
	var reqJson = ReqJsonRpc{
		"notice",
		params,
	}
	sender.Send(ws, reqJson)
}

func (sender *SenderStruct) Send(ws *websocket.Conn, reqJson interface{}) {
	s, err_json_marshal := json.Marshal(reqJson)
	if err_json_marshal != nil {
		// log
	}
	if err := websocket.Message.Send(ws, string(s)); err != nil {
		// log
	}
}
