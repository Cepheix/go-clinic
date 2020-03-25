package common

import "strings"

type Configuration struct {
	Server  Server
	Logging Logging
}

type Server struct {
	Address         string
	Port            string
	SwaggerDocument string
	SwaggerPrefix   string
}

type Logging struct {
	Development bool
}

func (server Server) FullAddress() string {
	return server.Address + ":" + server.Port
}

func (server Server) SwaggerDocumentAddress(prefix string) string {
	values := []string{"http:/", server.FullAddress(), prefix, server.SwaggerPrefix, server.SwaggerDocument}
	return strings.Join(values, "/")
}
