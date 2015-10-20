package controllers

import (
	"encoding/json"
	// "errors"
	"fmt"
	"github.com/realint/dbgutil"
	"golang.org/x/net/websocket"
	"net/http"
	"project/myim"
	"sixbyte/framework"
	"strconv"
)

type WebsocketController struct{}

type ReqJsonRpc struct {
	Method string            `json:"method"`
	Params map[string]string `json:"params"`
}

var Websocket = &WebsocketController{}

func (wsc *WebsocketController) Client(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t := framework.ViewTemplate("/websocket/client.html")
		t.Execute(w, nil)
	} else {
	}
}

func (wsc *WebsocketController) Server(ws *websocket.Conn) {

	defer func() {
		var reqJson ReqJsonRpc
		wsc.Logout(ws, reqJson)
	}()

	for {
		// 消息监听接收, 验证
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("无法接收.") // response and close
			break
		}
		var reqJson ReqJsonRpc
		if err := json.Unmarshal([]byte(reply), &reqJson); err != nil {
			fmt.Println("响应: ", err.Error()) // response and don't close
		}
		dbgutil.FormatDisplay("reqJson", reqJson)

		// 方法对应的操作
		if reqJson.Method == "login" {
			wsc.Login(ws, reqJson)
		}
		if reqJson.Method == "logout" {
			wsc.Logout(ws, reqJson)
		}

		if reqJson.Method == "sendto" {
			wsc.Sendto(ws, reqJson)
		}
	}
}

func (wsc *WebsocketController) Login(ws *websocket.Conn, reqJson ReqJsonRpc) {
	name := reqJson.Params["name"]
	if len(name) == 0 {
		myim.Sender.Error(ws, "1", "name must required")
		return
	}
	player := myim.NewPlayer(name, ws)
	player.LogIn()
}

func (wsc *WebsocketController) Logout(ws *websocket.Conn, reqJson ReqJsonRpc) {
	player := myim.FindPlayer(ws)
	if player == nil {
		// 没有这个用户
		return
	}
	player.LogOut()
}

func (wsc *WebsocketController) Sendto(ws *websocket.Conn, reqJson ReqJsonRpc) {
	player := myim.FindPlayer(ws)
	if player == nil {
		myim.Sender.Error(ws, "2", "请先登录.....")
		return
	}
	if reqJson.Params["content"] == "" || reqJson.Params["player_id"] == "" {
		myim.Sender.Error(ws, "1", "参数错误~")
		return
	}
	if reqJson.Params["player_id"] == "all" {
		player.SendToPlayers(reqJson.Params["content"])
		return
	}
	player_id, err_strconv_int := strconv.ParseInt(reqJson.Params["player_id"], 10, 64)
	if err_strconv_int != nil {
		myim.Sender.Error(ws, "1", "参数错误~")
		return
	}
	to := myim.FindPlayer(player_id)
	player.SendToPlayer(reqJson.Params["content"], to)
}
