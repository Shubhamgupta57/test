package main

import (
	"fmt"
	"integrations/cmd/server"
)

func main() {

	// r, err := api.InitHandlers()
	// if err != nil {
	// 	panic(err)
	// }
	// err = server.Server(r)
	// if err != nil {
	// 	panic(err)
	// }

	server, err := server.NewServer()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	server.StartServer()
	defer server.StopServer()
}
