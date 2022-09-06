package main

import (
	"leadster/oidc"
	"leadster/server"
)

func main() {
	server := server.New()
	oidc.New(server)
	<-server.Ctx.Done()
}
