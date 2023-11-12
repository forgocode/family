package web_im

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/webservice/service/web_im"
)

func ReceiveClientComm(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		HandshakeTimeout: time.Second * 10,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		newlog.Logger.Errorf("failed to generate websocket connection, err:%+v\n", err)
		return
	}
	// defer c.Close()
	uid := ctx.Request.Header.Get("uid")
	if uid == "" {
		newlog.Logger.Errorf("failed ot get uuid from header\n")
		return
	}
	web_im.AddWebSocketClient(uid, c)
}
