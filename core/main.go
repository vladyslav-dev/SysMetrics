package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemStats struct {
	CPU        CPUStats     `json:"cpu"`
	Memory     MemoryStats  `json:"memory"`
	Disk       DiskStats    `json:"disk"`
	Network    NetworkStats `json:"network"`
	SystemInfo SystemInfo   `json:"system_info"`
}

type CPUStats struct {
	UsagePercent float64 `json:"usage_percent"`
	Cores        int     `json:"cores"`
}

type MemoryStats struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskStats struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
	Free  uint64 `json:"free"`
}

type NetworkStats struct {
	UploadSpeed   uint64 `json:"upload_speed"`
	DownloadSpeed uint64 `json:"download_speed"`
}

type SystemInfo struct {
	OS       string `json:"os"`
	Uptime   string `json:"uptime"`
	Hostname string `json:"hostname"`
}

func main() {
	for {
		stats := getSystemStats()
		output, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(output))

		time.Sleep(1 * time.Second)
	}
}

func getSystemStats() SystemStats {
	return SystemStats{
		CPU:        GetCPUStats(),
		Memory:     GetMemoryStats(),
		Disk:       GetDiscStats(),
		Network:    GetNetworkStats(),
		SystemInfo: GetSystemInfo(),
	}
}

func GetCPUStats() CPUStats {
	usage, _ := cpu.Percent(0, false)
	cores, _ := cpu.Counts(false)
	return CPUStats{
		UsagePercent: usage[0],
		Cores:        cores,
	}
}

func GetMemoryStats() MemoryStats {
	vm, _ := mem.VirtualMemory()
	return MemoryStats{
		Total:       vm.Total,
		Used:        vm.Used,
		Free:        vm.Free,
		UsedPercent: vm.UsedPercent,
	}
}

func GetDiscStats() DiskStats {
	usage, _ := disk.Usage("/")
	return DiskStats{
		Total: usage.Total,
		Used:  usage.Used,
		Free:  usage.Free,
	}
}

func GetNetworkStats() NetworkStats {
	ioCounters, _ := net.IOCounters(false)
	if len(ioCounters) == 0 {
		return NetworkStats{}
	}
	return NetworkStats{
		UploadSpeed:   ioCounters[0].BytesSent,
		DownloadSpeed: ioCounters[0].BytesRecv,
	}
}

func GetSystemInfo() SystemInfo {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return SystemInfo{
		OS:       runtime.GOOS,
		Uptime:   fmt.Sprintf("%v hours", time.Since(time.Now().Add(-time.Duration(time.Now().Unix())*time.Second)).Hours()),
		Hostname: hostname,
	}
}
