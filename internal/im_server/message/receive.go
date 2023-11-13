package message

import (
	"fmt"

	"github.com/forgocode/family/internal/pkg/typed"
)

var msgChan = make(chan *typed.MessageInfo)

func ReceiveMessage() {
	for msg := range msgChan {
		//开始事务？
		//send message to db
		//find toUUID client
		//send
		fmt.Println(msg)
	}
}
