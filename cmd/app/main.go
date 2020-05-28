package main

import (
	"github.com/reynaldoqs/x_resolver/internal/infrastructure/server"
)

//const defaultPot = "8080"

func main() {
	//port
	/*	log.Println("starting API cmd")
		port := os.Getenv("PORT")
		if port == ""{
			port = defaultPort
		}
	*/

	// change start
	server.RegisterRouter()
}
