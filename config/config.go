package config

import (
	"flag"
)

type Config struct {
	ServerAddress                               string
	ServerConnectionMode                        string
	TimeoutInSeconds                            int
	WaitTimeUntilServerConnectionRetryInSeconds int
	WaitTimeUntilMessageReadRetryInSeconds      int
	Version                                     string
}

const AppVersion = "0.0.1"

func NewConfig() *Config {

	// Command-line flags
	serverAddress := flag.String("ip", DefaultServerAddress, "IP address to bind the server")
	port := flag.String("port", DefaultServerAddressPort, "Port to bind the server")
	serverConnectionMode := flag.String("mode", ServerConnectionModeSocket, "Server mode (socket or http)")

	// Parse the arguments
	flag.Parse()

	fullServerAddress := *serverAddress + ":" + *port

	return &Config{
		ServerAddress:        fullServerAddress,
		ServerConnectionMode: *serverConnectionMode,
		TimeoutInSeconds:     DefaultTimeoutInSeconds,
		WaitTimeUntilServerConnectionRetryInSeconds: 60 * DefaulWaitTimeUntilServerConnectionRetryInMinutes,
		WaitTimeUntilMessageReadRetryInSeconds:      60 * DefaultWaitTimeUntilMessageReadRetryInMinutes,
		Version:                                     AppVersion,
	}
}
