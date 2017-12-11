package main

import (
	"os"
	"fmt"
	"flag"
	"github.com/ageweile/pingpongclient/client"
	"github.com/ageweile/pingpongclient/server"
)

func main() {

	serverAddr := flag.String("s", "", "Server URL. (Required)")
	flag.Parse()

	fmt.Println("server: ", *serverAddr)

	if *serverAddr == "" {
		flag.PrintDefaults()
		os.Exit(1)

	}

	c := client.Client{}
	c.Login(*serverAddr) // initial login

	s := server.Server{}
	s.CreateRoutes() // register routes and start server
}
