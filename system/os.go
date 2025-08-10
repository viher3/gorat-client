package system

import (
	"io/ioutil"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func getUptime() string {
	if runtime.GOOS == "windows" {
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
	// Linux
	data, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return "unknown"
	}
	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return "unknown"
	}
	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return "unknown"
	}
	days := int(seconds) / 86400
	hours := (int(seconds) % 86400) / 3600
	minutes := (int(seconds) % 3600) / 60
	return strconv.Itoa(days) + " days " + strconv.Itoa(hours) + " hours " + strconv.Itoa(minutes) + " minutes"
}

func getKernelVersion() string {
	if runtime.GOOS == "windows" {
		out, err := exec.Command("cmd", "/C", "ver").Output()
		if err != nil {
			return "unknown"
		}
		return strings.TrimSpace(string(out))
	}
	// Linux
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func getCPUModel() string {
	if runtime.GOOS == "windows" {
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
	// Linux
	data, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return "unknown"
}

func getTotalMemory() string {
	if runtime.GOOS == "windows" {
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
	// Linux
	data, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				kb, err := strconv.Atoi(parts[1])
				if err != nil {
					return "unknown"
				}
				gb := float64(kb) / 1048576 // 1 GB = 1048576 kB
				return strconv.FormatFloat(gb, 'f', 2, 64) + " GB"
			}
		}
	}
	return "unknown"
}
