package main

import (
	"runtime"

	"github.com/josephbateh/go-api/server"
)

func main() {
	runtime.GOMAXPROCS(2)
	server.Start()
}
