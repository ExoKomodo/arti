package main

import (
	"leadster/oidc"
	"leadster/server"
)

func main() {
	server.Config()
	server := server.New()
	oidc.New(server)
	<-server.Ctx.Done()
}
