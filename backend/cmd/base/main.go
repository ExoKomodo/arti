package main

import (
	"leadster/server"
)

func main() {
	server := server.New()
	<-server.Ctx.Done()
}
