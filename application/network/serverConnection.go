package network

import (
	"fmt"

	"github.com/viher3/gorat-client/network"
)

func Execute() {

	privateIp, privateIpErr := network.GetPrivateIP()
	publicIp, publicIpErr := network.GetPublicIP()

	if privateIpErr != nil {
		fmt.Println("Error getting Private IP:", privateIpErr)
		return
	}

	if publicIpErr != nil {
		fmt.Println("Error getting Public IP:", publicIpErr)
		return
	}

	fmt.Println("")
	fmt.Println("Private IP:", privateIp)
	fmt.Println("Public IP:", publicIp)
	fmt.Println("")
}
