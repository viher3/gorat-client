//go:build linux
// +build linux

package system

import (
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
)

func getUptime() string {
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
	out, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

func getCPUModel() string {
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
