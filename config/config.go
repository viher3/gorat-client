package config

type Config struct {
	ServerAddress               string
	ServerConnectionMode        string
	TimeoutInSeconds            int
	WaitTimeUntilRetryInSeconds int
	Version                     string
}

const AppVersion = "0.0.1"

func NewConfig() *Config {
	return &Config{
		ServerAddress:               "192.168.1.0:1234",
		ServerConnectionMode:        ServerConnectionModeWebsocket,
		TimeoutInSeconds:            DefaultTimeoutInSeconds,
		WaitTimeUntilRetryInSeconds: 60 * 60 * DefaulWaitTimeUntilRetryInMinutes,
		Version:                     AppVersion,
	}
}
