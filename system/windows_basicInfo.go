//go:build windows
// +build windows

package system

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func getUptime() string {
	out, err := exec.Command("wmic", "os", "get", "LastBootUpTime", "/value").Output()
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(out), "\n")
	var bootTimeStr string
	for _, line := range lines {
		if strings.HasPrefix(line, "LastBootUpTime=") {
			bootTimeStr = strings.TrimPrefix(line, "LastBootUpTime=")
			bootTimeStr = strings.TrimSpace(bootTimeStr)
			break
		}
	}
	if bootTimeStr == "" {
		return "unknown"
	}
	// WMIC returns format: YYYYMMDDHHMMSS.milliseconds+timezone
	layout := "20060102150405"
	bootTimeStr = bootTimeStr[:14]
	bootTime, err := time.Parse(layout, bootTimeStr)
	if err != nil {
		return "unknown"
	}
	uptime := time.Since(bootTime)
	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60
	return strconv.Itoa(days) + " days " + strconv.Itoa(hours) + " hours " + strconv.Itoa(minutes) + " minutes"
}

func getKernelVersion() string {
	out, err := exec.Command("cmd", "/C", "ver").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func getCPUModel() string {
	out, err := exec.Command("wmic", "cpu", "get", "Name").Output()
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		return strings.TrimSpace(lines[1])
	}
	return "unknown"
}

func getTotalMemory() string {
	out, err := exec.Command("wmic", "ComputerSystem", "get", "TotalPhysicalMemory").Output()
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		memStr := strings.TrimSpace(lines[1])
		memBytes, err := strconv.ParseFloat(memStr, 64)
		if err != nil {
			return "unknown"
		}
		gb := memBytes / (1024 * 1024 * 1024)
		return strconv.FormatFloat(gb, 'f', 2, 64) + " GB"
	}
	return "unknown"
}
