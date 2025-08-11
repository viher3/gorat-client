package config

import (
	"flag"
)

type Config struct {
	ServerAddress               string
	ServerConnectionMode        string
	TimeoutInSeconds            int
	WaitTimeUntilRetryInSeconds int
	Version                     string
}

const AppVersion = "0.0.1"

func NewConfig() *Config {

	// Command-line flags
	serverAddress := flag.String("ip", DefaultServerAddress, "IP address to bind the server")
	port := flag.String("port", DefaultServerAddressPort, "Port to bind the server")
	serverConnectionMode := flag.String("mode", ServerConnectionModeWebsocket, "Server mode (websocket or http)")

	// Parse the arguments
	flag.Parse()

	fullServerAddress := *serverAddress + ":" + *port

	return &Config{
		ServerAddress:               fullServerAddress,
		ServerConnectionMode:        *serverConnectionMode,
		TimeoutInSeconds:            DefaultTimeoutInSeconds,
		WaitTimeUntilRetryInSeconds: 60 * 60 * DefaulWaitTimeUntilRetryInMinutes,
		Version:                     AppVersion,
	}
}
