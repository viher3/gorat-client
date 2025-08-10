package network

import (
	"fmt"
)

func Execute() {

	privateIp, privateIpErr := GetPrivateIP()
	publicIp, publicIpErr := GetPublicIP()

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
