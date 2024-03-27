package chat

import (
	"testing"
	"time"
)

func TestPluginRun(t *testing.T) {
	p := &ChatPlugin{
		ExecPath:   "/root/goWorkspace/family/bin/log_service",
		PluginName: "log_Service",
		ListenPort: 10002,
	}
	p.Run()

	time.Sleep(50 * time.Second)
}
