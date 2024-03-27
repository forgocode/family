package group

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/forgocode/family/internal/webservice/router/base"
)

type GroupPlugin struct {
	PluginName  string `json:"name" gorm:"column:name"`
	Md5         string `json:"md5" gorm:"column:md5"`
	Version     string `json:"version" gorm:"column:version"`
	Author      string `json:"author" gorm:"author"`
	Description string `json:"description" gorm:"description"`
	Status      string `json:"status" gorm:"column:status"`
	ExecPath    string `json:"execPath"`
	ListenPort  int32
}

func init() {
	p := &GroupPlugin{
		PluginName:  "群组服务",
		Version:     "0.0.1_base",
		Author:      "forgocode",
		Description: "激活群组功能，进行群聊",
		ExecPath:    "",
		ListenPort:  10002,
	}
	base.RegisterPlugin(p)
}

func (p *GroupPlugin) Name() string {
	return p.PluginName
}

func (p *GroupPlugin) Run() (*exec.Cmd, error) {
	cmd := exec.Command(p.ExecPath, "-port", strconv.Itoa(int(p.ListenPort)))
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Printf("failed to wait plugin: %+v, err: %+v\n", p, err)
		}
	}()
	return cmd, nil
}

func (p *GroupPlugin) Router() []base.RouterInfo {
	return []base.RouterInfo{}
}

func (p *GroupPlugin) Uninstall() {

}

func (p *GroupPlugin) Upgrade() {

}
