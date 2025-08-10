package config

type Config struct {
	CcServerAddress             string
	DefaultTimeoutInSeconds     int
	WaitTimeUntilRetryInSeconds int
	Version                     string
}

func NewConfig() *Config {
	return &Config{
		CcServerAddress:             "192.168.1.0:1234",
		DefaultTimeoutInSeconds:     30,
		WaitTimeUntilRetryInSeconds: 60 * 60 * 5,
		Version:                     "0.0.1",
	}
}
