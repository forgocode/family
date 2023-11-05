package sendlog

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	proto "github.com/forgocode/family/internal/grpcserver/proto/log"
	"github.com/forgocode/family/internal/pkg/newlog"
)

type msgStruct struct {
	ModuleCN string
	ModuleEN string
	MsgCN    string
	MsgEN    string
}

const (
	LoginCode = iota + 100000
	NewCategory
	DeleteCategory

	NewTag
	DeleteTag
)

const (
	SystemModuleEN = "system"
	SystemModuleCN = "系统模块"

	TagEN      = "Tags"
	TagCN      = "标签模块"
	CategoryEN = "Category"
	CategoryCN = "类别模块"
)

// 前三位模块，后三位递增
var msgMap = map[int32]msgStruct{
	LoginCode: {
		ModuleCN: SystemModuleCN,
		ModuleEN: SystemModuleEN,
		MsgCN:    "用户登录成功",
		MsgEN:    "user login successfully!",
	},
	NewCategory: {
		ModuleCN: CategoryCN,
		ModuleEN: CategoryEN,
		MsgCN:    "新建分类[%s]",
		MsgEN:    "new category [%s]",
	},
	DeleteCategory: {
		ModuleCN: CategoryCN,
		ModuleEN: CategoryEN,
		MsgCN:    "删除分类[%s]",
		MsgEN:    "delete category [%s]",
	},
	NewTag: {
		ModuleCN: TagCN,
		ModuleEN: TagEN,
		MsgCN:    "新建标签[%s]",
		MsgEN:    "new category [%s]",
	},
	DeleteTag: {
		ModuleCN: TagCN,
		ModuleEN: TagEN,
		MsgCN:    "删除标签[%s]",
		MsgEN:    "delete tag [%s]",
	},
}

func getModuleByLangAndCode(lang string, msgCode int32) string {
	if lang == "en" {
		return msgMap[msgCode].ModuleEN
	}
	return msgMap[msgCode].ModuleCN
}

func getMessageByLangAndCode(lang string, msgCode int32, detail ...string) string {
	if lang == "en" {
		if len(detail) != 0 {
			return fmt.Sprintf(msgMap[msgCode].MsgEN, detail)
		}
		return msgMap[msgCode].MsgEN
	}
	if len(detail) != 0 {
		return fmt.Sprintf(msgMap[msgCode].MsgCN, detail)
	}
	return msgMap[msgCode].MsgCN
}

func SendOperationLog(userID, lang string, msgCode int32, detail ...string) error {
	logInfo := &proto.OperationLogInfo{
		User:       userID,
		Module:     getModuleByLangAndCode(lang, msgCode),
		CreateTime: time.Now().UnixMilli(),
		Msg:        getMessageByLangAndCode(lang, msgCode, detail...),
	}
	return sendToGrpcServer(logInfo)

}

func sendToGrpcServer(info *proto.OperationLogInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conn, err := grpc.Dial("127.0.0.1:10000", grpc.WithInsecure())
	if err != nil {
		newlog.Logger.Errorf("failed to get conn from endpoint[%s], err info: %+v\n", "127.0.0.1:10000", err)
		return err
	}
	defer conn.Close()
	client := proto.NewOperationLogClient(conn)
	_, err = client.AddOperationLog(ctx, info)
	if err != nil {
		return err
	}
	return nil
}
