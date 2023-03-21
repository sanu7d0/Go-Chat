package main

import "client"

func main() {
	c := client.NewClient("localhost", "8000")
	c.Run()
}
