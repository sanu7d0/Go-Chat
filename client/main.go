package main

import (
	"fmt"
)

func main() {
	if err := TryLogin("Bob701"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Logged in")

	StartChat()
}
