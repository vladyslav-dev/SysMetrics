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

const (
	UnitPercentage = "%"
	UnitCores      = "Cores"
	UnitGB         = "GB"
	UnitMB         = "MB"
	UnitBytes      = "Bytes"
	UnitMBps       = "MBps"
	UnitHours      = "Hours"
)

type SystemMetrics struct {
	CPU        CPUMetrics     `json:"cpu"`
	Memory     MemoryMetrics  `json:"memory"`
	Disk       DiskMetrics    `json:"disk"`
	Network    NetworkMetrics `json:"network"`
	SystemInfo SystemInfo     `json:"system_info"`
}

type Metric struct {
	Original uint64 `json:"original"`
	Value    string `json:"value"`
	Unit     string `json:"unit"`
}

type CPUMetrics struct {
	UsagePercent Metric `json:"usage_percent"`
	Cores        Metric `json:"cores"`
}

type MemoryMetrics struct {
	Total       Metric `json:"total"`
	Used        Metric `json:"used"`
	Free        Metric `json:"free"`
	UsedPercent Metric `json:"used_percent"`
}

type DiskMetrics struct {
	Total Metric `json:"total"`
	Used  Metric `json:"used"`
	Free  Metric `json:"free"`
}

type NetworkMetrics struct {
	UploadSpeed   Metric `json:"upload_speed"`
	DownloadSpeed Metric `json:"download_speed"`
}

type SystemInfo struct {
	OS       string `json:"os"`
	Uptime   string `json:"uptime"`
	Hostname string `json:"hostname"`
}

func main() {
	for {
		stats := getSystemMetrics()
		output, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(output))

		time.Sleep(1 * time.Second)
	}
}

func getSystemMetrics() SystemMetrics {
	return SystemMetrics{
		CPU:        GetCPUMetrics(),
		Memory:     GetMemoryMetrics(),
		Disk:       GetDiskMetrics(),
		Network:    GetNetworkMetrics(),
		SystemInfo: GetSystemInfo(),
	}
}

func GetCPUMetrics() CPUMetrics {
	usage, _ := cpu.Percent(0, false)
	cores, _ := cpu.Counts(false)
	return CPUMetrics{
		UsagePercent: Metric{
			Original: uint64(usage[0]),
			Value:    fmt.Sprintf("%.2f", usage[0]),
			Unit:     UnitPercentage,
		},
		Cores: Metric{
			Original: uint64(cores),
			Value:    fmt.Sprintf("%d", cores),
			Unit:     UnitCores,
		},
	}
}

func GetMemoryMetrics() MemoryMetrics {
	vm, _ := mem.VirtualMemory()
	return MemoryMetrics{
		Total: Metric{
			Original: vm.Total,
			Value:    formatBytes(vm.Total),
			Unit:     UnitGB,
		},
		Used: Metric{
			Original: vm.Used,
			Value:    formatBytes(vm.Used),
			Unit:     UnitGB,
		},
		Free: Metric{
			Original: vm.Free,
			Value:    formatBytes(vm.Free),
			Unit:     UnitGB,
		},
		UsedPercent: Metric{
			Original: uint64(vm.UsedPercent),
			Value:    fmt.Sprintf("%.2f", vm.UsedPercent),
			Unit:     UnitPercentage,
		},
	}
}

func GetDiskMetrics() DiskMetrics {
	usage, _ := disk.Usage("/")
	return DiskMetrics{
		Total: Metric{
			Original: usage.Total,
			Value:    formatBytes(usage.Total),
			Unit:     UnitGB,
		},
		Used: Metric{
			Original: usage.Used,
			Value:    formatBytes(usage.Used),
			Unit:     UnitGB,
		},
		Free: Metric{
			Original: usage.Free,
			Value:    formatBytes(usage.Free),
			Unit:     UnitGB,
		},
	}
}

func GetNetworkMetrics() NetworkMetrics {
	ioCounters, _ := net.IOCounters(false)
	if len(ioCounters) == 0 {
		return NetworkMetrics{}
	}
	return NetworkMetrics{
		UploadSpeed: Metric{
			Original: ioCounters[0].BytesSent,
			Value:    formatBytes(ioCounters[0].BytesSent),
			Unit:     UnitMBps,
		},
		DownloadSpeed: Metric{
			Original: ioCounters[0].BytesRecv,
			Value:    formatBytes(ioCounters[0].BytesRecv),
			Unit:     UnitMBps,
		},
	}
}

func GetSystemInfo() SystemInfo {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	uptime := time.Since(time.Now().Add(-time.Duration(time.Now().Unix()) * time.Second)).Hours()

	return SystemInfo{
		OS:       runtime.GOOS,
		Uptime:   fmt.Sprintf("%.2f %s", uptime, UnitHours),
		Hostname: hostname,
	}
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d %s", bytes, UnitBytes)
	}
	div, exp := unit, 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	value := float64(bytes) / float64(div)
	return fmt.Sprintf("%.2f", value)
}
