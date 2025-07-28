package main

import (
	"fmt"

	"github.com/viher3/gorat-client/application/network"
	"github.com/viher3/gorat-client/config"
)

func main() {

	fmt.Println("############################")
	fmt.Println("### goRat client v", config.Version, "###")
	fmt.Println("############################")

	network.Execute()
}
