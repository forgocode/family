package typed

import (
	"time"

	"github.com/gorilla/websocket"
)

type MessageInfo struct {
	FromUID     string
	FromName    string
	ToUID       string
	Type        MessageType
	Context     string
	GroupID     string
	ContextType int
	Time        time.Time
}

type MessageType int

const (
	BroadCast MessageType = iota + 1
)

type WebSocketClient struct {
	Client   *websocket.Conn
	UserName string
}
