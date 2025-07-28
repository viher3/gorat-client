package network

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var publicIPServices = []string{
	"https://api.ipify.org",
	"https://ifconfig.me/ip",
	"https://ipinfo.io/ip",
}

func GetPrivateIP() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			// Solo IPv4 privadas
			if ipv4 := ip.To4(); ipv4 != nil && isPrivateIP(ipv4) {
				return ipv4.String(), nil
			}
		}
	}
	return "", nil
}

func GetPublicIP() (string, error) {
	client := http.Client{Timeout: 30 * time.Second}

	for _, url := range publicIPServices {
		resp, err := client.Get(url)
		if err != nil {
			continue // try next service
		}
		defer resp.Body.Close()

		ip, err := io.ReadAll(resp.Body)
		if err != nil {
			continue // try next service
		}
		if resp.StatusCode == http.StatusOK && len(ip) > 0 {
			return string(ip), nil
		}
	}
	return "", fmt.Errorf("Can't get public IP from external services")
}

func isPrivateIP(ip net.IP) bool {
	// 10.0.0.0/8
	if ip[0] == 10 {
		return true
	}
	// 172.16.0.0/12
	if ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31 {
		return true
	}
	// 192.168.0.0/16
	if ip[0] == 192 && ip[1] == 168 {
		return true
	}
	return false
}
