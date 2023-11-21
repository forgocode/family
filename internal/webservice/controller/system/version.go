package system

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"github.com/forgocode/family/internal/conf"
	"github.com/forgocode/family/internal/pkg/newlog"
	"github.com/forgocode/family/internal/pkg/response"
)

func GetVersion(ctx *gin.Context) {
	info := conf.GetConfig().Base
	response.Success(ctx, info, 1)
}

func GetMonitor(ctx *gin.Context) {
	type info struct {
		CpuCount    int     `json:"cpuCount"`
		CpuPercent  float64 `json:"cpuPercent"`
		UsedMemory  uint64  `json:"usedMemory"`
		TotalMemory uint64  `json:"totalMemory"`
		UsedDisk    uint64  `json:"usedDisk"`
		TotalDisk   uint64  `json:"totalDisk"`
		RunningTime int     `json:"runningTime"`
	}
	totalMemory, usedMemory := getMemoryInfo()
	totalDisk, usedDisk := getDiskInfo()

	result := &info{
		CpuCount:    runtime.NumCPU(),
		CpuPercent:  getCpuPercent(),
		UsedMemory:  usedMemory / 1024 / 1024 / 1024,
		TotalMemory: totalMemory / 1024 / 1024 / 1024,
		UsedDisk:    usedDisk / 1024 / 1024 / 1024,
		TotalDisk:   totalDisk / 1024 / 1024 / 1024,
		RunningTime: 0,
	}
	response.Success(ctx, result, 1)

}

func getCpuPercent() float64 {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		newlog.Logger.Errorf("failed to get cpu percent, err:%+v\n", err)
		return 0
	}
	return percent[0]
}

func getMemoryInfo() (uint64, uint64) {
	info, err := mem.VirtualMemory()
	if err != nil {
		newlog.Logger.Errorf("failed to get memory info, err:%+v\n", err)
		return 0, 0
	}
	return info.Total, info.Used
}

func getDiskInfo() (uint64, uint64) {
	info, err := disk.Usage("/")
	if err != nil {
		newlog.Logger.Errorf("failed to get disk info, err:%+v\n", err)
		return 0, 0
	}

	return info.Total, info.Used
}
