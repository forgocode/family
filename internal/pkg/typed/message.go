package typed

import (
	"time"

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
	Time        time.Time   `json:"time"`
}

type MessageType int

const (
	BroadCast MessageType = iota + 1
)

type WebSocketClient struct {
	Client   *websocket.Conn
	UserName string
}
