package main

import (
	"fmt"

	"github.com/viher3/gorat/network"
)

func main() {
	ip, err := network.GetPrivateIP()
	if err != nil {
		fmt.Println("Error al obtener IP privada:", err)
		return
	}
	fmt.Println("IP privada actual:", ip)
}
