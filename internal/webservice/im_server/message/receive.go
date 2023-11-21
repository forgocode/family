package message

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/typed"
)

var msgChan = make(chan *typed.MessageInfo, 1024)

func SetMessage2Chan(msg *typed.MessageInfo) error {
	select {
	case msgChan <- msg:
		fmt.Printf("%+v\n", msg)
		return nil
	default:
		return errors.New("channal is full, please try later")
	}
}

func ReceiveMessage() {
	for msg := range msgChan {
		data, err := json.Marshal(msg)
		if err != nil {
			continue
		}

		switch msg.Type {
		case typed.SystemBroadCast:
			broadCastMessage(data)
		default:
			err := sendMsgToUser(msg.ToUID, data)
			if err != nil {
				newlog.Logger.Errorf("failed to send msg to")
			}

		}
		//向外发送日志记录

	}
}
