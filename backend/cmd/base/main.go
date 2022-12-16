package main

import (
	"arti/oidc"
	"arti/server"
	"fmt"
)

func main() {
	server.Config()
	server, err := server.New()
	if err != nil {
		fmt.Printf("failed to create server: %x\n", err)
	}
	oidc.New(*server)
	<-server.Ctx.Done()
}
