package main

import (
	"github.com/fschr/go/server"
)

func main() {
	s := server.NewServer(":3000")
	s.Run()
}
