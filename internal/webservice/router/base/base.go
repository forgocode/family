package base

import (
	"os/exec"
	"sync"

	"github.com/gin-gonic/gin"
)

type Plugin interface {
	Run() (*exec.Cmd, error)
	Uninstall()
	Upgrade()
	Name() string
	Router() []RouterInfo
}

type RouterInfo struct {
	Group      string
	Path       string
	Method     string
	Handles    []gin.HandlerFunc
	Middleware []gin.HandlerFunc
}

const (
	Running = iota + 1
	Initing
	Stopped
	Upgrading
)

func RegisterPlugin(p Plugin) {
	pluginManagerCenter.RegisterPlugin(p)
}

func RegisterRouter(engine *gin.Engine) {
	// 注册插件管理的路由
	pluginManagerCenter.RegisterRouter(engine)

}

func (m *pluginManager) RegisterRouter(engine *gin.Engine) {
	pluginManagerCenter.mu.Lock()
	defer pluginManagerCenter.mu.Unlock()
	for p, _ := range pluginManagerCenter.items {
		// if s.status != Running {
		// 	continue
		// }
		for _, r := range p.Router() {
			switch r.Method {
			case "GET":
				engine.Group(r.Group).Use(r.Middleware...).GET(r.Path, r.Handles...)
			case "POST":
				engine.Group(r.Group).Use(r.Middleware...).POST(r.Path, r.Handles...)
			case "PUT":
				engine.Group(r.Group).Use(r.Middleware...).PUT(r.Path, r.Handles...)
			case "DELETE":
				engine.Group(r.Group).Use(r.Middleware...).DELETE(r.Path, r.Handles...)
			}
		}
	}
}

type pluginManager struct {
	mu    sync.RWMutex
	items map[Plugin]PluginInfo
}

type PluginInfo struct {
	cmdInfo *exec.Cmd
	status  int
}

var pluginManagerCenter = &pluginManager{
	mu:    sync.RWMutex{},
	items: make(map[Plugin]PluginInfo),
}

func (m *pluginManager) RegisterPlugin(p Plugin) {
	m.mu.Lock()
	defer m.mu.Unlock()
	info := PluginInfo{
		cmdInfo: nil,
		status:  Initing,
	}

	m.items[p] = info
}

func (m *pluginManager) Load(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.items {
		if k.Name() != name {
			continue
		}
		cmd, err := k.Run()
		if err != nil {
			return err
		}
		v.cmdInfo = cmd
		v.status = Running
		m.items[k] = v
	}
	return nil
}

func (m *pluginManager) UnLoad(name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.items {
		if k.Name() != name {
			continue
		}
		if v.status != Running {
			continue
		}
		//TODO: ?可以取消吗
		return v.cmdInfo.Cancel()
	}
	return nil
}
