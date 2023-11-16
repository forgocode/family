package message

import (
	"github.com/forgocode/family/internal/pkg/newlog"

	client_manager "github.com/forgocode/family/internal/webservice/im_server/client"
)

func broadCastMessage(data []byte) {
	clients := client_manager.ListClient()
	for _, c := range clients {
		err := c.Client.WriteMessage(0, data)
		newlog.Logger.Debugf("server write message info: %+v\n", err)
		if err != nil {
			newlog.Logger.Errorf("write: %+v\n", err)
		}
	}
}

func sendMsgToUser(uid string, data []byte) error {
	toC, err := client_manager.FindClientByUid(uid)
	err = toC.Client.WriteMessage(1, data)
	if err != nil {
		newlog.Logger.Errorf("write: %+v\n", err)
		return err
	}
	return nil
}
