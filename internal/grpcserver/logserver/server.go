package logserver

import (
	"context"
	"errors"
	"net"
	"time"

	"google.golang.org/grpc"

	logproto "github.com/forgocode/family/internal/grpcserver/proto/log"
	messageproto "github.com/forgocode/family/internal/grpcserver/proto/message"
	stationnoticeproto "github.com/forgocode/family/internal/grpcserver/proto/station_notice"
	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/webservice/database/mongo"
	"github.com/forgocode/family/internal/webservice/model"
	"github.com/forgocode/family/pkg/uuid"
)

type Server struct {
	logproto.UnimplementedOperationLogServer
}
type stationNoticeServer struct {
	stationnoticeproto.UnimplementedHandleStationNoticeServer
}

type messageServer struct {
	messageproto.UnimplementedHandleMessageServer
}

var operationLog = make(chan *model.OperationLogInfo, 512)

func (s *Server) AddOperationLog(ctx context.Context, msg *logproto.OperationLogInfo) (*logproto.Response, error) {
	newlog.Logger.Debugf("receive log info: %+v\n", msg)
	log := &model.OperationLogInfo{}
	log.UUID = uuid.GetUUID()
	log.Convert(msg)
	select {
	case operationLog <- log:
		return &logproto.Response{
			Status: 200,
		}, nil
	default:
		return &logproto.Response{Status: 200}, errors.New("grpc server can't receive operation log\n")
	}
}

func Start() {
	go handleOperation()
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		newlog.Logger.Errorf("failed to listen 10000 port: %+v\n", err)
		return
	}
	s := grpc.NewServer()
	registerGrpc(s)
	err = s.Serve(listener)
	if err != nil {
		newlog.Logger.Errorf("failed to server grpc server, err: %+v\n", err)
		return
	}
}

func registerGrpc(s *grpc.Server) {
	logproto.RegisterOperationLogServer(s, &Server{})
	messageproto.RegisterHandleMessageServer(s, *&messageServer{})
	stationnoticeproto.RegisterHandleStationNoticeServer(s, &stationNoticeServer{})
}

func handleOperation() {
	for {
		select {
		case msg := <-operationLog:
			c, err := mongo.GetMongoClient("operation_log")
			if err != nil {
				newlog.Logger.Errorf("failed to get mongo client, err: %+v\n", err)
				continue
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

			_, err = c.InsertOne(ctx, msg)
			if err != nil {
				newlog.Logger.Errorf("failed to insert log info: %+v\n", err)
			}
			cancel()
		}
	}
}
