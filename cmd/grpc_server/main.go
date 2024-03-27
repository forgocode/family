package main

import (
	"flag"
	"os"

	"github.com/forgocode/family/internal/grpcserver/logserver"
	"github.com/forgocode/family/internal/pkg/newlog"
)

var port uint64

func main() {
	flag.Uint64Var(&port, "port", 10000, "the port for cmd listen")
	flag.Parse()
	newlog.InitLogger("", os.Stdout)
	newlog.Logger.Infof("grpc server is start succfully, listen port: %d!\n", port)

	//数据库
	//从kafka读消息，或者作为一个grpc server接受其他模块的消息
	//启动
	logserver.Start(port)
	select {}
}

type Handler interface {
	AddLog()
	UpdateLog()
}
