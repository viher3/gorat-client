package main

import (
	"fmt"

	"github.com/viher3/gorat-client/config"
	"github.com/viher3/gorat-client/system"
)

func main() {
	conf := config.NewConfig()
	fmt.Println("############################")
	fmt.Println("### goRat client v", conf.Version, "###")
	fmt.Println("############################")

	systemInfo := system.GetFullInfo()
	if systemInfo == nil {
		fmt.Println("Failed to retrieve system information.")
		return
	}

	fmt.Println("System Information:")
	fmt.Println("OS:", systemInfo["os"])
	fmt.Println("Architecture:", systemInfo["arch"])
	fmt.Println("Hostname:", systemInfo["hostname"])
	fmt.Println("User:", systemInfo["user"])
	fmt.Println("Uptime:", systemInfo["uptime"])
	fmt.Println("Kernel Version:", systemInfo["kernel"])
	fmt.Println("CPU Model:", systemInfo["cpu"])
	fmt.Println("Memory:", systemInfo["memory"])
	fmt.Println("Public IP:", systemInfo["publicIp"])
	fmt.Println("Private IP:", systemInfo["privateIp"])
}
