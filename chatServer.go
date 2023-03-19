package main

import (
	"server"
)

func main() {
	s := server.NewServer("8000")
	s.Run()
}
