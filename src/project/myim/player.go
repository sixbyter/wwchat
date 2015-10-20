package myim

import (
	"errors"
	"golang.org/x/net/websocket"
	"time"
)

type Player struct {
	Id     int64           `json:"id"`
	Name   string          `json:"name"`
	client *websocket.Conn `json:"-"`
}

var (
	Players = []*Player{}
)

func NewPlayer(name string, client *websocket.Conn) *Player {
	unix := time.Now().Unix()
	var player = &Player{
		unix,
		name,
		client,
	}
	Players = append(Players, player)
	return player
}

func FindPlayer(what interface{}) *Player {
	switch what.(type) {
	case int64:
		for _, v := range Players {
			if v.Id == what {
				return v
			}
		}
	case *websocket.Conn:
		for _, v := range Players {
			if v.client == what {
				return v
			}
		}
	default:
		return nil
	}
	return nil
}

/**
 * 用户登录 = NewPlayer
 */
func (player *Player) LogIn() {
	player.UpdatePlayerList()
	// 通知其他用户 go 方法
	go player.noticeLogIn()
}

/**
 * 用户登出
 */
func (player *Player) LogOut() error {
	for k, v := range Players {
		if v.Id == player.Id {
			Players = append(Players[:k], Players[k+1:]...)
			player.client.Close()
			// 通知其他用户
			go player.noticeLogOut()
			return nil
		}
	}
	return errors.New("不存在此用户.")
}

/**
 * 通知其他用户我登录了
 * @param  {[type]} player *Player)      noticeLogIn( [description]
 * @return {[type]}        [description]
 */
func (player *Player) noticeLogIn() {
	for _, v := range Players {
		if v.Id == player.Id {
			continue
		}
		Sender.Notice(v.client, player.Name+" 登录了.")
		v.UpdatePlayerList()
	}
}

/**
 * 通知其他用户我退出了
 * @param  {[type]} player *Player)      noticeLogOut( [description]
 * @return {[type]}        [description]
 */
func (player *Player) noticeLogOut() {
	for _, v := range Players {
		if v.Id == player.Id {
			continue
		}
		Sender.Notice(v.client, player.Name+" 退出了.")
		v.UpdatePlayerList()
	}
}

func (player *Player) UpdatePlayerList() {
	type PlayersListReqJsonRpc struct {
		Method string               `json:"method"`
		Params map[string][]*Player `json:"params"`
	}
	var params = make(map[string][]*Player)
	params["players"] = Players
	var reqJson = PlayersListReqJsonRpc{
		"UpdatePlayerList",
		params,
	}
	Sender.Send(player.client, reqJson)
}

/**
 * 用户发送信息给特定的人
 */
func (player *Player) SendToPlayer(content string, to *Player) {
	player.SendMessage("SendToPlayer", content, to)
}

func (player *Player) SendToPlayers(content string) {
	for _, v := range Players {
		if v.Id == player.Id {
			continue
		}
		player.SendMessage("SendToPlayers", content, v)
	}
}

func (player *Player) SendMessage(method, content string, to *Player) {
	type Message struct {
		From    *Player `json:"from"`
		To      *Player `json:"to"`
		Content string  `json:"content"`
		Time    int64   `json:"time"`
	}
	type MessageReqJsonRpc struct {
		Method string              `json:"method"`
		Params map[string]*Message `json:"params"`
	}
	unix := time.Now().Unix()
	var params = make(map[string]*Message)
	message := &Message{
		player,
		to,
		content,
		unix,
	}
	params["message"] = message
	var reqJson = MessageReqJsonRpc{
		method,
		params,
	}
	Sender.Send(to.client, reqJson)
}
