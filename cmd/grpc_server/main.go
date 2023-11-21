package main

import (
	"os"

	"github.com/forgocode/family/internal/grpcserver/logserver"
	"github.com/forgocode/family/internal/pkg/newlog"
)

func main() {
	newlog.InitLogger("", os.Stdout)
	newlog.Logger.Infof("log server is start!\n")

	//数据库
	//从kafka读消息，或者作为一个grpc server接受其他模块的消息
	//启动
	logserver.Start()
	newlog.Logger.Infof("grpc server is start succfully!\n")
	select {}
}

type Handler interface {
	AddLog()
	UpdateLog()
}
