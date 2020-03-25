package common

type Configuration struct {
	Server  Server
	Logging Logging
}

type Server struct {
	Address string
}

type Logging struct {
	Development bool
}
