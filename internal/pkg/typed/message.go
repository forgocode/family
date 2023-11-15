package typed

import (
	"github.com/gorilla/websocket"
)

type MessageInfo struct {
	FromUID     string      `json:"fromUID"`
	FromName    string      `json:"fromName"`
	ToUID       string      `json:"toUID"`
	Type        MessageType `json:"type"`
	Context     string      `json:"context"`
	GroupID     string      `json:"groupID"`
	ContextType int         `json:"contextType"`
	Time        int64       `json:"time"`
}

type MessageType int

const (
	//系统广播消息不存mongo
	SystemBroadCast MessageType = iota + 1
)

type WebSocketClient struct {
	Client   *websocket.Conn
	UserName string
}
