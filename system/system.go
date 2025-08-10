package system

import (
	"os"
	"os/user"
	"runtime"

	"github.com/viher3/gorat-client/network"
)

func GetBasicInfo() map[string]string {
	u, _ := user.Current()

	info := make(map[string]string)
	info["os"] = runtime.GOOS
	info["arch"] = runtime.GOARCH
	info["hostname"], _ = os.Hostname()
	info["user"] = u.Username
	info["uptime"] = getUptime()
	info["kernel"] = getKernelVersion()
	info["cpu"] = getCPUModel()
	info["memory"] = getTotalMemory()
	return info
}

func GetNetworkInfo() map[string]string {
	networkInfo := make(map[string]string)
	privateIP, _ := network.GetPrivateIP()
	publicIP, _ := network.GetPublicIP()

	networkInfo["privateIp"] = privateIP
	networkInfo["publicIp"] = publicIP
	return networkInfo
}

func GetFullInfo() map[string]string {
	result := make(map[string]string)
	for k, v := range GetBasicInfo() {
		result[k] = v
	}
	for k, v := range GetNetworkInfo() {
		result[k] = v
	}
	return result
}
