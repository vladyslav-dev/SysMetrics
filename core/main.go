package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

const (
	UnitPercentage = "%"
	UnitCores      = "Cores"
	UnitGB         = "GB"
	UnitMbps       = "Mbps"
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
	usage, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("Error fetching CPU usage: %v", err)
	}

	var usagePercent float64
	if len(usage) > 0 {
		usagePercent = usage[0]
	} else {
		usagePercent = 0
	}

	cores, err := cpu.Counts(false)
	if err != nil {
		log.Printf("Error fetching CPU core count: %v", err)
		cores = 0
	}

	return CPUMetrics{
		UsagePercent: Metric{
			Original: uint64(usagePercent),
			Value:    fmt.Sprintf("%.2f", usagePercent),
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
			Value:    formatBytes(vm.Total, UnitGB),
			Unit:     UnitGB,
		},
		Used: Metric{
			Original: vm.Used,
			Value:    formatBytes(vm.Used, UnitGB),
			Unit:     UnitGB,
		},
		Free: Metric{
			Original: vm.Free,
			Value:    formatBytes(vm.Free, UnitGB),
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
			Value:    formatBytes(usage.Total, UnitGB),
			Unit:     UnitGB,
		},
		Used: Metric{
			Original: usage.Used,
			Value:    formatBytes(usage.Used, UnitGB),
			Unit:     UnitGB,
		},
		Free: Metric{
			Original: usage.Free,
			Value:    formatBytes(usage.Free, UnitGB),
			Unit:     UnitGB,
		},
	}
}

var (
	prevBytesSent uint64
	prevBytesRecv uint64
)

func GetNetworkMetrics() NetworkMetrics {
	ioCounters, _ := net.IOCounters(false)
	if len(ioCounters) == 0 {
		return NetworkMetrics{}
	}

	currentBytesSent := ioCounters[0].BytesSent
	currentBytesRecv := ioCounters[0].BytesRecv

	uploadDelta := currentBytesSent - prevBytesSent
	downloadDelta := currentBytesRecv - prevBytesRecv

	prevBytesSent = currentBytesSent
	prevBytesRecv = currentBytesRecv

	return NetworkMetrics{
		UploadSpeed: Metric{
			Original: uploadDelta,
			Value:    formatBytes(uploadDelta, UnitMbps),
			Unit:     UnitMbps,
		},
		DownloadSpeed: Metric{
			Original: downloadDelta,
			Value:    formatBytes(downloadDelta, UnitMbps),
			Unit:     UnitMbps,
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

func formatBytes(bytes uint64, unit string) string {
	switch unit {
	case UnitGB:
		return fmt.Sprintf("%.2f", float64(bytes)/(1024*1024*1024))
	case UnitMbps:
		return fmt.Sprintf("%.2f", (float64(bytes)*8)/(1024*1024))
	default:
		return fmt.Sprintf("%d", bytes)
	}
}
