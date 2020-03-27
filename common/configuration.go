package common

import "strings"

type Configuration struct {
	Server   Server
	Logging  Logging
	Database Database
}

type Server struct {
	Address         string
	Port            string
	SwaggerDocument string
	SwaggerPrefix   string
}

func (server Server) FullAddress() string {
	return server.Address + ":" + server.Port
}

func (server Server) SwaggerDocumentAddress() string {
	values := []string{"http:/", server.FullAddress(), server.SwaggerPrefix, server.SwaggerDocument}
	return strings.Join(values, "/")
}

type Logging struct {
	Development bool
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	Ssl      string
}

func (databaseConfiguration Database) ConnectionString() string {
	host := "host=" + databaseConfiguration.Host
	port := "port=" + databaseConfiguration.Port
	user := "user=" + databaseConfiguration.User
	password := "password=" + databaseConfiguration.Password
	dbName := "dbname=" + databaseConfiguration.DbName
	ssl := "sslmode=" + databaseConfiguration.Ssl

	elements := []string{host, port, user, password, dbName, ssl}

	return strings.Join(elements, " ")
}
