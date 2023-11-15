package web_im

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/typed"
	"github.com/forgocode/family/internal/webservice/database/redis"
	client_manager "github.com/forgocode/family/internal/webservice/im_server/client"
)

const offLineDuration = 30

func AddWebSocketClient(uid, userName string, c *websocket.Conn) {
	rs, err := redis.GetRedisClient()
	if err != nil {
		newlog.Logger.Errorf("failed to get redis client, err: %+v\n", err)
		// return
	}
	client := &typed.WebSocketClient{
		Client:   c,
		UserName: "",
	}
	//没找到client，不存在发送广播消息
	if rs.Get(uid).Err() != nil {
		//发送广播消息
		msg := &typed.MessageInfo{
			FromUID:     "",
			FromName:    "",
			ToUID:       "",
			Type:        typed.SystemBroadCast,
			Context:     fmt.Sprintf("用户%s上线", userName),
			GroupID:     "",
			ContextType: 0,
			Time:        time.Now().UnixMilli(),
		}
		data, err := json.Marshal(msg)
		if err != nil {
			newlog.Logger.Errorf("failed to marshal message, err: %+v\n", err)
			return
		}
		clients := client_manager.ListClient()
		for _, c := range clients {
			err = c.Client.WriteMessage(0, data)
			newlog.Logger.Debugf("server write message info: %+v\n", err)
			if err != nil {
				newlog.Logger.Errorf("write: %+v\n", err)
			}
		}

		go func() {
			for {
				mt, message, err := c.ReadMessage()
				if err != nil {
					newlog.Logger.Errorf("read: %+v\n", err)
					break
				}
				fmt.Printf("%+v\n", string(message))
				msg := &typed.MessageInfo{}
				err = json.Unmarshal(message, msg)
				if err != nil {
					newlog.Logger.Errorf("failed to Unmarshal message, err: %+v\n", err)
					continue
				}
				newlog.Logger.Debugf("server receive message info: %+v\n", msg)
				toC, err := client_manager.FindClientByUid(msg.ToUID)

				//todo 放到队列中
				err = toC.Client.WriteMessage(mt, data)
				newlog.Logger.Debugf("server write message info: %+v\n", msg)
				//立即投递，放入消息队列，缓存
				if err != nil {
					newlog.Logger.Errorf("write: %+v\n", err)
					break
				}
			}
		}()
	}
	client_manager.AddClient(uid, client)

}
