package http_server

import "flag"

var (
	Addr = flag.String("http", ":8080", "http listen address")
)