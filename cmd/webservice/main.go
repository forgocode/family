package main

import (
	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/webservice/database"
	"github.com/forgocode/family/internal/webservice/im_server/message"
	"github.com/forgocode/family/internal/webservice/router"
	_ "net/http/pprof"
	"os"
)

func main() {
	newlog.InitLogger("", os.Stdout)
	newlog.Logger.Infof("Family Community System is started!\n")
	database.Start()
	go message.ReceiveMessage()
	router.Start()

}
