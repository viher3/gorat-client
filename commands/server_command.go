package commands

type ServerCommand interface {
	action() string
	payload() string
}
