package main

import "github.com/ChrisLo0751/xap-server/internal/server"

func main() {
	http := server.NewHttpServer()
	// Listen and Server in 0.0.0.0:8080
	http.Run()
}
