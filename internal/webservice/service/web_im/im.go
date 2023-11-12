package web_im

import (
	"encoding/json"

	"github.com/gorilla/websocket"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/typed"
	"github.com/forgocode/family/internal/webservice/database/redis"
)

const offLineDuration = 15

func AddWebSocketClient(uid string, c *websocket.Conn) {
	rs, err := redis.GetRedisClient()
	if err != nil {
		return
	}
	client := typed.WebSocketClient{
		Client:   c,
		UserName: "",
	}
	//没找到client，不存在发送广播消息
	if rs.Get(uid).Err() != nil {
		//发送广播消息
		//启动goroutine来接收消息
		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					newlog.Logger.Errorf("read:", err)
					break
				}
				msg := &typed.MessageInfo{}
				err = json.Unmarshal(message, msg)
				if err != nil {
					newlog.Logger.Errorf("failed to Unmarshal message, err: %+v\n", err)
					continue
				}
				newlog.Logger.Debugf("server receive message info: %+v\n", msg)
				//放入消息队列
				if err != nil {
					newlog.Logger.Errorf("write:", err)
					break
				}
			}
		}()
	}
	if rs.Set(uid, client, 15*offLineDuration).Err() != nil {
		return
	}

}
