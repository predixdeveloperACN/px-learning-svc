package main

import (
	serv "github.com/predixdeveloperACN/px-learning-svc/server"
)

func main() {
	serv.SetupServer()
	serv.StartServer()
}
